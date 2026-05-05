package main

import (
	"fmt"
)

func handleInit() error {
	db, err := getDB()
	if err != nil {
		return fmt.Errorf("failed to get database: %w", err)
	}
	defer db.Close()

	createTablesSQL := `
	CREATE TABLE IF NOT EXISTS categories (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE
		);
		
		CREATE TABLE IF NOT EXISTS packages (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			category_id INTEGER NOT NULL,
			notes TEXT,
			FOREIGN KEY (category_id) REFERENCES categories(id)
			);`

	_, err = db.Exec(createTablesSQL)
	if err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}

	fmt.Println("テーブルの初期化が完了しました。")
	return nil
}
