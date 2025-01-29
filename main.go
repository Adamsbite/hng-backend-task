package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Response struct {
	Email          string `json:"email"`
	CurrentDateTime string `json:"current_datetime"`
	GitHubURL      string `json:"github_url"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers krk34m
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Construct the response
	response := Response{
		Email:          "your-email@example.com",
		CurrentDateTime: time.Now().UTC().Format(time.RFC3339),
		GitHubURL:      "https://github.com/yourusername/hng-backend-task",
	}

	// Write the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", handler)

	// Start the server
	port := "8080"
	println("Server running on port", port)
	http.ListenAndServe(":"+port, nil)
}
