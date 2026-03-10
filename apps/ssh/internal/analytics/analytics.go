package analytics

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"charm.land/log/v2"
)

var (
	url     string
	webID   string
	enabled bool
	client  = &http.Client{Timeout: 5 * time.Second}
)

func Init() {
	url = os.Getenv("UMAMI_URL")
	webID = os.Getenv("UMAMI_WEBSITE_ID")
	enabled = url != "" && webID != ""

	if enabled {
		log.Info("Analytics enabled", "url", url)
	}
}

type payload struct {
	Type    string `json:"type"`
	Payload event  `json:"payload"`
}

type event struct {
	Website  string `json:"website"`
	Hostname string `json:"hostname"`
	URL      string `json:"url"`
	Title    string `json:"title"`
	Language string `json:"language"`
	Screen   string `json:"screen"`
	Referrer string `json:"referrer"`
}

func TrackVisitor(username, remoteAddr string) {
	if !enabled {
		return
	}

	hostname := os.Getenv("UMAMI_HOSTNAME")
	if hostname == "" {
		hostname = "farhanaulianda.my.id"
	}

	body, _ := json.Marshal(payload{
		Type: "event",
		Payload: event{
			Website:  webID,
			Hostname: hostname,
			URL:      "/ssh",
			Title:    "SSH Session - " + username,
			Language: "en-US",
			Screen:   "1920x1080",
			Referrer: "",
		},
	})

	go func() {
		req, err := http.NewRequest("POST", url+"/api/send", bytes.NewReader(body))
		if err != nil {
			return
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36")

		resp, err := client.Do(req)
		if err != nil {
			log.Error("Analytics failed", "error", err)
			return
		}
		defer resp.Body.Close()
		respBody, _ := io.ReadAll(resp.Body)
		log.Info("Analytics sent", "status", resp.StatusCode, "response", string(respBody))
	}()
}
