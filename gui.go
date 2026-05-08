package main

import (
	"fmt"
	"net/http"
)

func handleGUI() error {
	mux := http.NewServeMux()


	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return fmt.Errorf("GUIサーバーの起動に失敗しました: %w", err)
	}
	return nil
}

type ErrorResponse struct {
	Message string `json:"message"`
}
