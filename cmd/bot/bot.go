package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	redis "gopkg.in/redis.v3"

	"github.com/bwmarrin/discordgo"
)

var (
	buffer = make([][]byte, 0)
	r      *redis.Client
)

type Play struct {
	GuildID   string
	ChannelID string
}

func main() {
	var (
		Token = flag.String("t", "", "Discord Authentication Token")
		Redis = flag.String("r", "", "Redis connection string")
	)
	flag.Parse()

	err := loadSound()
	if err != nil {
		fmt.Println("error loading sound: ", err)
		return
	}

	if *Redis != "" {
		r = redis.NewClient(&redis.Options{Addr: *Redis, DB: 0})

		_, err := r.Ping().Result()
		if err != nil {
			fmt.Println("failed to connect to redis", err)
			return
		}
	}

	discord, err := discordgo.New("Bot " + *Token)
	if err != nil {
		fmt.Println("error creating Discord session", err)
	}

	discord.AddHandler(messageCreate)

	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection", err)
		return
	}

	fmt.Println("bot is running.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println("message created " + m.Content)

	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "!shame") || strings.HasPrefix(m.Content, "!shamebell") {
		s.ChannelMessageSend(m.ChannelID, ":bell: Shame!")

		// Find channel
		c, err := s.State.Channel(m.ChannelID)
		if err != nil {
			// Couldn't find the channel.
			return
		}

		// Find guild for channel.
		g, err := s.State.Guild(c.GuildID)
		if err != nil {
			// Couldn't find guild.
			return
		}

		// Look for message sender in guild's voice states.
		for _, vs := range g.VoiceStates {
			if vs.UserID == m.Author.ID {

				err = playSound(s, g.ID, vs.ChannelID)
				if err != nil {
					fmt.Println("error playing sound", err)
				}
				return
			}
		}

	}
}

func loadSound() error {

	file, err := os.Open("./assets/shamebell.dca")
	if err != nil {
		fmt.Println("Error opening dca file :", err)
		return err
	}

	var opuslen int16

	for {
		// Read opus frame length from dca file.
		err = binary.Read(file, binary.LittleEndian, &opuslen)

		// If this is the end of the file, just return.
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			err := file.Close()
			if err != nil {
				return err
			}
			return nil
		}

		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return err
		}

		// Read encoded pcm from dca file.
		InBuf := make([]byte, opuslen)
		err = binary.Read(file, binary.LittleEndian, &InBuf)

		// Should not be any end of file errors
		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return err
		}

		// Append encoded pcm data to the buffer.
		buffer = append(buffer, InBuf)
	}
}

func playSound(s *discordgo.Session, guildID, channelID string) error {
	fmt.Println("playing sound")

	// Join the channel.
	vc, err := s.ChannelVoiceJoin(guildID, channelID, false, true)
	if err != nil {
		return err
	}

	// Track stats in redis.
	play := &Play{
		GuildID:   guildID,
		ChannelID: channelID,
	}
	go trackStats(play)

	// Sleep for specified amount of time before playing sound.
	time.Sleep(250 * time.Millisecond)

	// Start speaking.
	vc.Speaking(true)

	// Send buffer data.
	for _, buff := range buffer {
		vc.OpusSend <- buff
	}

	// Stop speaking.
	vc.Speaking(false)

	// Sleep for specified time.
	time.Sleep(250 * time.Millisecond)

	// Leave voice channel.
	vc.Disconnect()

	return nil
}

func trackStats(play *Play) {
	if r == nil {
		return
	}

	_, err := r.Pipelined(func(pipe *redis.Pipeline) error {
		pipe.Incr("shamebell:total")
		pipe.SAdd("shamebell:guilds", play.GuildID)
		pipe.SAdd("shamebell:channels", play.ChannelID)
		return nil
	})

	if err != nil {
		fmt.Println("Failed to track stats in redis")
	}
}
