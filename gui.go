package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"time"
)

//go:embed frontend/dist/*
var frontend embed.FS

func handleGUI() error {
	mux := http.NewServeMux()

	subFS, err := fs.Sub(frontend, "frontend/dist")
	mux.Handle("/", http.FileServer(http.FS(subFS)))
	mux.HandleFunc("/api/getpackages", getPackages)

	server := &http.Server{
		Addr:              "127.0.0.1:8080",
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
