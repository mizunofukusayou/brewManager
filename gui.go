package main

import (
	"fmt"
	"net/http"
	"time"
)

func handleGUI() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/getpackages", getPackages)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       15 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("GUIサーバーの起動に失敗しました: %w", err)
	}
	return nil
}

type ErrorResponse struct {
	Message string `json:"message"`
}
