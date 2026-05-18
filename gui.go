package main

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"time"
)

//go:embed frontend/dist/*
var frontend embed.FS

type api struct {
	db *sql.DB
}

func handleGUI() error {
	db, err := getDB()
	if err != nil {
		return fmt.Errorf("データベースの読み込みに失敗しました: %w", err)
	}
	defer db.Close()
	a := &api{db: db}

	mux := http.NewServeMux()

	subFS, err := fs.Sub(frontend, "frontend/dist")
	mux.Handle("GET /", http.FileServer(http.FS(subFS)))
	mux.HandleFunc("GET /api/getpackages", a.getPackages)

	server := &http.Server{
		Addr:              "127.0.0.1:8080",
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       15 * time.Second,
	}

	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		return fmt.Errorf("ポートの確保に失敗しました: %w", err)
	}
	fmt.Println("👀 access http://127.0.0.1:8080")
	err = server.Serve(ln)
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("サーバーで予期せぬエラーが発生しました: %w", err)
	}
	if err == http.ErrServerClosed {
		return fmt.Errorf("サーバーが正常終了しました: %w", err)
	}
	return nil
}

type ErrorResponse struct {
	Message string `json:"message"`
}
