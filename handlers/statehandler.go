package handlers

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/data"
	"net/http"
	"sync"
)

func StateHandler(w http.ResponseWriter, r *http.Request, mu *sync.RWMutex, currentState *data.State) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	for {
		mu.Lock()
		state := currentState
		mu.Unlock()

		var stateName string
		switch *state {
		case data.Loading:
			stateName = "Loading"
		case data.Success:
			stateName = "Success"
		case data.Error:
			stateName = "Error"
		}

		// Stream the state as a JSON object
		payload := map[string]string{"state": stateName}
		data, _ := json.Marshal(payload)
		fmt.Fprintf(w, "data: %s\n\n", data)
		w.(http.Flusher).Flush()

		
	}
}
