BOT_BINARY=bot
WEB_BINARY=web

.PHONY: all
all: bot web

bot: cmd/bot/bot.go
	go build -o ${BOT_BINARY} cmd/bot/bot.go

web: cmd/webserver/web.go static
	go build -o ${WEB_BINARY} cmd/web/web.go

npm: static/package.json
	cd static && npm install .

.PHONY: static
static: npm gulp

.PHONY: clean
clean:
	rm -r ${BOT_BINARY} ${WEB_BINARY} static/build/