package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type response struct {
	Slack_name   string    `json:"slack_name"`
	Current_time time.Time `json:"utc_time"`
	Current_day  string    `json:"current_day"`
	Track        string    `json:"track"`
	Github_file  string    `json:"github_file_url"`
	Github_repo  string    `json:"github_repo_url"`
	Status_code  int       `json:"status_code"`
}

func (res *response) Endpoint(w http.ResponseWriter, r *http.Request) {

	res.Current_time = time.Now().UTC()
	res.Current_day = time.Now().Weekday().String()
	res.Track = "backend"
	res.Github_file = os.Getenv("GITHUB_FILE_URL")
	res.Github_repo = os.Getenv("GITHUB_REPO_URL")
	res.Status_code = 200

	url_values := r.URL.Query()
	slack_name := url_values.Get("slack_name")
	track := url_values.Get("track")

	if slack_name == "" {
		res.Slack_name = os.Getenv("SLACK_NAME")
	} else {
		res.Slack_name = slack_name
	}

	if track == "" {
		res.Track = "backend"
	} else {
		res.Track = track
	}

	json.NewEncoder(w).Encode(res)

}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading environment files...", err)
	}

	handler := &response{}

	port := os.Getenv("PORT")
	mux := http.NewServeMux()
	mux.HandleFunc("/api", handler.Endpoint)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())

}
