package web_status

import (
	"GoRecordurbate/modules/bot"
	"GoRecordurbate/modules/config"
	"GoRecordurbate/modules/file"
	"encoding/json"
	"net/http"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}

	type Rec struct {
		Status string `json:"status"`
	}

	var recorder Rec
	config.Reload(file.Config_json_path, &config.Streamers)
	// Assuming config.Settings.App.Streamers is accessible
	for _, s := range bot.Bot.ListRecorders() {
		if s.IsRecording {
			recorder.Status = "Running"
			break // No need to continue checking
		} else {
			recorder.Status = "Stopped"

		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recorder)
}
