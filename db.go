package main

import (
	"database/sql"
	"fmt"
)

func getDB() (*sql.DB, error) {
	// データベースファイルを開く（存在しない場合は作成される）
	db, err := sql.Open("sqlite", "brewmanager.db?_pragma=foreign_keys(1)")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// 接続確認
	err = db.Ping()
	if err != nil {
		closeErr := db.Close()
		if closeErr != nil {
			return nil, fmt.Errorf("failed to ping database: %w (additionally failed to close database: %v)", err, closeErr)
		}
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	return db, nil
}

type Category struct {
	ID   int
	Name string
}

type Package struct {
	ID         int
	Name       string
	CategoryID int
	Notes      string
}
