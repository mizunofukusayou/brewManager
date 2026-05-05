package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	_ "modernc.org/sqlite"
)

var usageError = errors.New("usage error")

func usageErrorMessage(err error) string {
	msg := err.Error()
	const prefix = "usage error:"
	if strings.HasPrefix(msg, prefix) {
		return strings.TrimSpace(strings.TrimPrefix(msg, prefix))
	}
	return msg
}

func main() {
	err := run()
	if err != nil {
		if errors.Is(err, usageError) {
			fmt.Fprintln(os.Stderr, usageErrorMessage(err))
			os.Exit(2)
		} else {
			log.Printf("Error: %v", err)
			os.Exit(1)
		}
	}
}

func run() error {
	args := os.Args
	if len(args) < 2 {
		return fmt.Errorf("%w:コマンドを指定してください。例: bm init", usageError)
	}

	command := args[1]
	switch command {
	case "init":
		err := handleInit()
		if err != nil {
			return fmt.Errorf("failed to initialize database: %w", err)
		}

	case "add":
		err := handleAdd(args)
		if err != nil {
			return fmt.Errorf("failed to add item: %w", err)
		}

	default:
		return fmt.Errorf("%w:%sは不明なコマンドです", usageError, args[1])
	}

	return nil
}
