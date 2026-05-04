package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func main() {
	// データベースファイルを開く（存在しない場合は作成される）
	db, err := sql.Open("sqlite", "brewmanager.db?_pragma=foreign_keys(1)")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 接続確認
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// 2. テーブル作成（リレーションシップあり）
	// まず categories を作り、次にそれを使う packages を作る
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
		log.Fatal(err)
	}

	fmt.Println("テーブルが作成されました。")
}
