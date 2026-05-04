package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

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

	args := os.Args
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "コマンドを指定してください。例: bm init")
		os.Exit(2)
	}

	command := args[1]
	switch command {
	case "init":
		// テーブル作成（リレーションシップあり）
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

	case "add":
		if len(args) < 3 {
			fmt.Fprintln(os.Stderr, "引数が足りません。例: bm add category <名前> または bm add package <パッケージ名> <カテゴリID>")
			os.Exit(2)
		}

		switch args[2] {
		case "category":
			if len(args) < 4 {
				fmt.Fprintln(os.Stderr, "引数が足りません。例: bm add category <名前>")
				os.Exit(2)
			}
			name := args[3]
			_, err = db.Exec("INSERT INTO categories (name) VALUES (?)", name)
			if err != nil {
				log.Println("failed to insert category:", err)
				os.Exit(1)
			}

		case "package":
			if len(args) < 5 {
				fmt.Fprintln(os.Stderr, "引数が足りません。例: bm add package <パッケージ名> <カテゴリID>")
				os.Exit(2)
			}
			name := args[3]
			categoryID, err := strconv.Atoi(args[4])
			if err != nil {
				fmt.Fprintln(os.Stderr, "カテゴリIDは数値で指定してください。")
				os.Exit(2)
			}
			_, err = db.Exec("INSERT INTO packages (name, category_id) VALUES (?, ?)", name, categoryID)
			if err != nil {
				log.Println("failed to insert package:", err)
				os.Exit(1)
			}

		default:
			fmt.Fprintln(os.Stderr, "不明なサブコマンドです:", args[2])
			os.Exit(2)
		}

	default:
		fmt.Fprintln(os.Stderr, "不明なコマンドです:", command)
		os.Exit(2)
	}
}
