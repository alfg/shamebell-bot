# SHAMEBELL
Shame your friends in Discord for doing stupid things. :bell:

![](/assets/profile.png)

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

## Deployment
TODO

## Docker
TODO

## Resources
* https://github.com/bwmarrin/discordgo - Discord Golang Bindings
* https://github.com/hammerandchisel/airhornbot - This project was modeled after airhornbot
* https://discordapp.com/developers - Discord Developers Portal
* https://github.com/facebookincubator/create-react-app - React frontend boilerplate

## License
MIT