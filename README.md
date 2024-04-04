# Helldivers 2 Bot
This discord bot will send campaign status updates
to the provided channels once an hour.
This is based off of an unoffical [api](https://helldiverstrainingmanual.com/api)
and may stop working at some point without notice.

## Usage
To use this bot start setup a JSON config file with keys 
`channels` and `token`.
`token` will contain you discord bot token and
`channels` will contain a list of channels to post to.
By default this binary will look for `./config.json`, 
this can be changed using the `-config` option if needed.

### Example 
```json
{
  "channels": [
    "<REPLACE_WITH_CHANNEL>"
  ],
  "token": "<REPLACE_WITH_TOKEN>"
}
```

### Docker 
Build with docker 
```
docker build -t helldivers2-bot . 
```

Run in container
```
docker run -dit --rm --name helldivers2-bot helldivers2-bot
```

Stop container 
```
docker kill helldivers2-bot
```
