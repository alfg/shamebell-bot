# SHAMEBELL
Shame your friends in Discord for doing stupid things. :bell:

![](/assets/profile.png)

[![Discord Bots](https://discordbots.org/api/widget/servers/350849408075825152.svg)](https://discordbots.org/bot/350849408075825152)

## Installation
Simply go to https://shamebellbot.com/ to install the bot on your Discord server!


## Usage
Just type `!shame` while in a voice channel!


## Develop
Shamebell-bot requires the following for development:
* Go 1.4+
* Redis

To get started, first create an app and bot on Discord:
https://discordapp.com/developers/applications

Then git clone and install Go dependencies:
```b
git clone https://github.com/alfg/shamebell-bot.git /go/src/github.com/alfg/shamebell-bot
go get -u github.com/golang/dep/cmd/dep
dep ensure
```

You can now run the bot and web backend in separate processes:
```
go run cmd/web/web.go -r "<redis host:port>"

go run cmd/bot/bot.go -t "<your discord bot token>" -r "<redis host:port>"
```

Now build the frontend:
```
cd static && yarn

# Run the development server
yarn start
```

You should now have 3 separate processes: The bot, web backend and web frontend.

Load `http://localhost:3000/` in your browser.


## Docker
A `docker-compose.yml` is provided to easily setup and launch the bot, web server, and redis instance.

Pre-built Docker images are available and tagged for use at https://hub.docker.com/r/alfg/shamebell-bot/tags.

This makes it easy to deploy if your server has Docker installed using `docker-compose` or `docker-swarm`.

* Add your bot's auth token to `docker-compose.yml`'s, bot entrypoint `-t` flag.
* Run `docker-compose`:
```
docker-compose up
```
* Load `http://localhost:4000/` in browser to see the website.

#### Updating Version
Update `docker-compose.yml` bot and web tag to the latest version, then run:
```
docker-compose up -d --no-deps
```


## Resources
* https://github.com/bwmarrin/discordgo - Discord Golang Bindings
* https://github.com/hammerandchisel/airhornbot - This project was modeled after airhornbot
* https://discordapp.com/developers - Discord Developers Portal
* https://github.com/facebookincubator/create-react-app - React frontend boilerplate

## License
MIT
