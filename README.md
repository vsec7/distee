# DisTee ( Discord Tee )

<p align="center">
<a href="https://goreportcard.com/report/github.com/vsec7/distee"><img src="https://goreportcard.com/badge/github.com/vsec7/distee"></a>
<a href="https://github.com/vsec7/distee/issues"><img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat"></a>
<a href="https://twitter.com/verry__d"><img src="https://img.shields.io/twitter/follow/verry__d.svg?logo=twitter"></a>
</p>

DisTee is a GO tool that works like tee command, Feed input to distee through the stdin.

Crafted By : github.com/vsec7


## Installation
```
go get -u github.com/vsec7/distee
```

## Setup Configuration
```
distee -setup
```

## Basic Usage :
```
 ▶ echo "Hello Cath" | distee
 ▶ cat file.txt | distee -c <channel_id>
 ▶ cat file.txt | distee -c <channel_id> -t <title> -code

Options :
  -setup, --setup                  Setup Configuration
  -c, --c <channel_id>             Send message to custom channel_id
  -t, --t <title>                  Send message with title
  -code, --code                    Send message with code markdown
  -config, --config <config.yaml>  Set custom config.yaml location
```

## General Questions

[?] How to find the token ? <a href="https://www.writebots.com/discord-bot-token/"> READ HERE </a>

[?] How to find the channel id ? 

https://discord.com/channels/8173952126616XXXXX/<8174070495026XXXXX> <= Channel ID

## Dependencies
- github.com/bwmarrin/discordgo
