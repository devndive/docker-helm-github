package main

import (
	"fmt"
	"net/http"

	"gopkg.in/ini.v1"
)

type configuration struct {
	stage      string
	backendUrl string
}

func (h *configuration) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<!DOCTYPE html>")
	fmt.Fprintf(w, "<html><head><title>Sample APP</title><body>")
	fmt.Fprintf(w, "<h2>Application<h2> <h3>My properties are:</h3><ul>")
	fmt.Fprintf(w, "<li>Stage: "+h.stage+"</li>")
	fmt.Fprintf(w, "<li>BackendURL: "+h.backendUrl+"</li>")
	fmt.Fprintf(w, "</ol>")
	fmt.Fprintf(w, "</body></html>")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Healthy\n")
}

func main() {
	cfg, err := ini.LooseLoad("settings.ini", "/config/settings.ini")
	if err != nil {
		fmt.Printf("Fail to read configuration file: %v\n", err)
	}

	config := configuration{}
	config.stage = cfg.Section("").Key("stage").String()
	config.backendUrl = cfg.Section("backend").Key("url").String()

	fmt.Println("Starting server on port :8080")
	http.Handle("/", &config)
	http.HandleFunc("/health", healthHandler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Kaputt 😟: %v", err)
	}
}
