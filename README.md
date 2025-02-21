# GoRecordurbate WebUI

Readme is incomplete and will modified a lot during development.

Checkout the [API doc](https://github.com/luna-nightbyte/GoRecordurbate-WebUI/blob/main/internal/docs/API_REFERENCE.md) for info on how to use the api.

## Intro
A [golang](https://go.dev/) version of recordurbate with some differences. One key difference is that this doesent use a deamon, but instead runs a webserver (Web UI). It can be compiled into a binary file and started as a service, docker container or whatever you prefer. 
GoRecordurbate send a request to check if the [spesific](https://github.com/luna-nightbyte/GoRecordurbate/blob/ec0b1fa79e2bb82cf948bef3415ace3aac52e523/modules/bot/bot.go#L176) user is online rather than requesting a [list of 500](https://github.com/luna-nightbyte/GoRecordurbate/blob/ec0b1fa79e2bb82cf948bef3415ace3aac52e523/modules/bot/bot.go#L175) and checking that for the correct user. 
### Usage
Default login (will be modified):
User: `admin`
Password: `password`

#### Build
Building the code wil create a binary for your os system. Golang is [cross-compatible](https://go.dev/wiki/GccgoCrossCompilation) for windows, linux and mac.
```bash
go mod init GoRecordurbate # Only run this line once
go mod tidy
go build
```
#### Source
```bash
go mod init GoRecordurbate # Only run this line once
go mod tidy
go run main.go
```

### Notes
This is un-tested on Windows and Mac, but golang is cross-compatible which means that this should run just as fine on Windows as on Linux.

A release will be made once i have finished fixing the bare minimum below:
- [x] Start recording
- [ ] Restart recording
- [x] stop recording
- [x] Add / delete streamer
- [ ] Import streamers
- [x] Export streamers
- [ ] Show log
- [ ] View videos in web ui
- [ ] Docker example
- [ ] Service example
- [ ] Embed index file into code

## WebUI (Will probably be modified)


<p align="center">
  <img src="https://github.com/user-attachments/assets/35e4633b-702b-45f9-9075-a8522a6b334b" alt="Login page"/>
  <img src="https://github.com/user-attachments/assets/56913a17-b200-4416-b32f-fe92460cc34f" alt="Streamers tab"/>
  <img src="https://github.com/user-attachments/assets/83186720-f056-41f1-b6fa-433c9869d9c1" alt="Control tab"/>
</p>

## Other

### Ideas, but not planned
- Log online time of streamers and save to csv for graph plotting. Can help understand the work-hours of different streamers.
- Option to login and use follower list instead of config?
- Option for max video length (and size?)





## Thanks

Special thanks to [oliverjrose99](https://github.com/oliverjrose99) for the inspiration and their work on [Recordurbate](https://github.com/oliverjrose99/Recordurbate)
