#!/bin/bash

# Gets guilds server count from redis and POSTs to DiscordBots stats.
# Get TOKEN from https://discordbots.org/bot/${USERID}/edit

USERID=
TOKEN=
COUNT="$(docker-compose run redis-cli redis-cli -h redis SMEMBERS "shamebell:guilds" | wc -l)"

curl "https://discordbots.org/api/bots/${USERID}/stats" \
    -i -v \
    -X POST \
    -H "Content-Type: application/json" \
    -H "Authorization: ${TOKEN}" \
    -d '{ "server_count": "'${COUNT}'" }'