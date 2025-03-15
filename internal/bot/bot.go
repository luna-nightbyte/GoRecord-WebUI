package bot

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"GoStreamRecord/internal/db"
	"GoStreamRecord/internal/recorder"
)

// Bot encapsulates the recording bot’s state.
// TODO: finish bot structure
type controller struct {
	mux        sync.Mutex
	status     []recorder.Recorder
	isFirstRun bool
	logger     *log.Logger
	// ctx is used to signal shutdown.
	ctx    context.Context
	cancel context.CancelFunc
}

var Bot *controller

func Init() *controller {
	Bot = NewBot(log.New(os.Stdout, "lpg.log", log.LstdFlags))
	return Bot
}

// NewBot creates a new Bot, sets up its cancellation context.
func NewBot(logger *log.Logger) *controller {
	ctx, cancel := context.WithCancel(context.Background())
	b := &controller{
		logger:     logger,
		ctx:        ctx,
		cancel:     cancel,
		status:     []recorder.Recorder{},
		isFirstRun: true,
	}
	return b
}

// RecordLoop starts the main loop for a given streamer.
// It checks for online status, starts recording if not already recording, and listens for a shutdown signal.
func (b *controller) RecordLoop(streamerName string) {
	// Always ensure config exists
	if err := recorder.WriteYoutubeDLdb(); err != nil {
		log.Println("Error writing youtube-dl db:", err)
		return
	}

	var wg sync.WaitGroup
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	// Loop over configured streamers.
	for i1 := range db.Config.Streamers.Streamers {
		configIndex := i1
		streamer := db.Config.Streamers.Streamers[configIndex]
		if streamer.Name == streamerName || streamerName == "" {
			// Start a new recorder if one isn’t already running.
			if b.isRecorderActive(streamer.Name) {
				fmt.Println("Recorder was active")
				continue
			}
			b.AddProcess(db.Config.Streamers.Streamers[i1].Provider, streamer.Name)

			// Find the Recorder for the streamer.
			for i2 := range b.status {
				// Ensure correct name is being used.
				streamer.Name = b.status[i2].Website.Interface.TrueName(streamer.Name)
				if b.status[i2].Website.Username != streamer.Name {
					continue
				}
				wg.Add(1)
				// Pass the index and streamer name into the closure to avoid capture issues.
				go func(status *recorder.Recorder, sName string) {
					defer wg.Done()
					stopStatus := false
					for {
						// Exit the goroutine if the bot is cancelled.
						select {
						case <-b.ctx.Done():
							return
						default:
						}

						if stopStatus {

							b.StopProcessIfRunning(status)
							log.Println("Stopped!")
							// If not a restart, exit.
							b.mux.Lock()
							if !status.IsRestarting {
								b.mux.Unlock()
								return
							}
							b.mux.Unlock()
							stopStatus = false
							status.IsRestarting = false
						} else {
							b.mux.Lock()
							if b.isRecorderActive(sName) {
								b.mux.Unlock()
								return
							}
							b.mux.Unlock()

							db.Read("settings", "settings.json", &db.Config.Settings)

							log.Printf("Checking %s online status...", sName)
							fmt.Println(status.Website.Interface)
							if !status.Website.Interface.IsOnline(sName) {
								log.Printf("Streamer %s is not online.", sName)
								return
							}
							log.Printf("Streamer %s is online!", sName)
							// Mark as recording.
							b.mux.Lock()
							status.IsRecording = true
							b.mux.Unlock()

							status.StartRecording(sName)

							b.mux.Lock()
							status.IsRecording = false
							status.StopSignal = true
							b.mux.Unlock()

							log.Printf("Recording for %s finished", sName)
							stopStatus = true
						}
						time.Sleep(time.Duration(db.Config.Settings.App.Loop_interval) * time.Minute)
					}
				}(&b.status[i2], streamer.Name)
			}
			if streamer.Name == streamerName {
				break
			}
		}
	}
	time.Sleep(time.Duration(db.Config.Settings.App.Loop_interval) * time.Minute)
	wg.Wait()
}
