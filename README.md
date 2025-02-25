# GoRecordurbate WebUI

__API NOTE__: The API is still open after adding adding login. The plan is to use browser cookies and optinally API secret key in the future for credentials via API. But every API endpoint will be reachable through a simple __curl__ command for now. This is not really any issue for anyone who only plan to use this inside their home network. But __Do NOT__ expose the port for this app on your router!! 

## Core Features
### Recorder:
- Streamer status checks bypasses rate limits; recording rate limits still under testing.
- Start, stop, and restart all recordings.
- View active recorders in real-time (More data will be added)
- Start, stop restart  individual recordings as needed.
- Add or delete streamers dynamically (no need to restart the recorder).
- Import/export streamer lists for easier management.

### WebUI:
- Secure login using cookies for each client (prevents unauthorized access).
- Manage multiple user accounts.
- Directly view logs and recorded videos through the WebUI.
- Watch live streams directly from the WebUI.
- Check streamer online status.
### Setup & Deployment:
- Docker configuration and service examples for straightforward deployment.
- Install [Golang](https://go.dev/doc/install) to run the source or build binary.
## Usage

|Username|Password|
|-|-|
|`admin`|see [this](https://github.com/luna-nightbyte/Recordurbate-WebUI/tree/main?tab=readme-ov-file#reset-password)|

### Setup
__important__: You will still need to have the `internal/settings` folder and it's content in the same folder structure when running this app. That means that you'll have to copy that along with any binary you build.

- Download this repo and open a terminal in this folder. Ask ChatGPT how to find the folder path and how to move into it via cli if you dont know.

- Copy [`.env.example`](https://github.com/luna-nightbyte/Recordurbate-WebUI/blob/main/.env.example) to `.env` and add your own session key. (se the one below as long as the app is not exposed outside your local network__IF__ you dont know how to create one. _hint: just ask chatGPT_). 
    - Can be generated with this command on linux.: `head -c 32 /dev/urandom | base64`
      I.e:
      ```bash
      user@user:~/Recordurbate$ head -c 32 /dev/urandom | base64
      # Output
      Fl60B6sTqAUARyDiC6GIor8AIu6QXLF2RMvWK1Wz3eE=
      ```

#### Optional config settings
The main settings can be found in [`settings.json`](https://github.com/luna-nightbyte/Recordurbate-WebUI/blob/main/internal/settings/settings.json):
```json
{
  "app": {
    "port": 8055,
    "loop_interval_in_minutes": 2,
    "video_output_folder": "output/videos",
    "rate_limit": {
      "enable": true,
      "time": 5
    },
    "default_export_location": "./output/list.txt"

  },
  "youtube-dl": {
    "binary": "youtube-dl"
  },
  "auto_reload_config": true
}
```
#### Reset password
To change forgotten password, start the program with the `reset-pwd` argument. I.e:
```
./GoRecordurbate reset-pwd admin newpassword 
```
New login for the user `admin` would then be `newpassword`
### Build
Building the code wil create a binary for your os system. Golang is [cross-compatible](https://go.dev/wiki/GccgoCrossCompilation) for windows, linux and mac.
```bash
go mod init GoRecordurbate # Only run this line once
go mod tidy
go build
./GoRecordurbate #windows will have 'GoRecordurbate.exe'
```
### Source
```bash
go mod init GoRecordurbate # Only run this line once
go mod tidy
go run main.go
```

## Notes
This is un-tested on Windows and Mac, but golang is cross-compatible which means that this should run just as fine on Windows as on Linux.

## WebUI (Will probably be modified)


<p align="center">
  <img src="https://github.com/user-attachments/assets/edf30517-de6a-4f91-9ab4-89f9c91d7779" alt="Login page"/>
  <img src="https://github.com/user-attachments/assets/5d939bc0-778b-42c8-a453-eb30c13e95e2" alt="Video tab"/>
  <img src="https://github.com/user-attachments/assets/0ce5b2c1-e7f3-47bb-96e9-1532915dd5e4" alt="individual tab"/>
  <img src="https://github.com/user-attachments/assets/7736fac5-5ce8-4634-8179-6ea2cf03969b" alt="User settings tab"/>
  
  <img src="https://github.com/user-attachments/assets/ced11119-8e74-4c15-8aff-6c31242f8fe5" alt="Streamers tab"/>
  <img src="https://github.com/user-attachments/assets/edc136e5-0238-463e-b8f3-d4b1b7e74687" alt="Livestream tab"/>
</p>

_Online status with a small bug at the time of uploading this.._
## Other

### Todo / Ideas
- Log online-time of streamers and save to csv for graph plotting. Can help understand the work-hours of different streamers.
- Select and delete videos
- Option to login to the streaming site and use follower list instead of config?
- Option for max video length (and size?)
- ~~headless mode without webui~~ (Abandoned because i will not create all the logic for handling the various arguments myself. Others can create a PR if they want to.)
- Move frontend to Vue
  - Btter for organizing components being re-used
- Build a default docker image
- Individual recorders in UI
  - Start/Restart individual recorders (in progress)
  - set max lenght/size (could be optional to use one of either)
  - view current recording length
  - view current recording video
- Better video view


### Disclaimer 
Unauthorized resale, redistribution, or sharing of recorded content that you do not own or have explicit permission to distribute is strictly prohibited. Users are solely responsible for ensuring compliance with all applicable copyright and privacy laws. The creator of this recorder assumes no liability for any misuse or legal consequences arising from user actions.

## Thanks

Special thanks to [oliverjrose99](https://github.com/oliverjrose99) for the initial inspiration and their work on [Recordurbate](https://github.com/oliverjrose99/Recordurbate). Initial code of this project was directly inspired by their project.
