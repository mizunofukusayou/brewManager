package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "modernc.org/sqlite"
)

var usageError = errors.New("usage error")

func main() {
	err := run()
	if err != nil {
		if errors.Is(err, usageError) {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}else {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
	}
}

func run() error {
	args := os.Args
	if len(args) < 2 {
		return fmt.Errorf("%w:コマンドを指定してください。例: bm init", usageError)
	}
	
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
			return fmt.Errorf("%w:引数が足りません。例: bm add category <名前> または bm add package <パッケージ名> <カテゴリID>", usageError)
		}
	
		switch args[2] {
		case "category":
			if len(args) < 4 {
				return fmt.Errorf("%w:引数が足りません。例: bm add category <名前>", usageError)
			}
			name := args[3]
			_, err = db.Exec("INSERT INTO categories (name) VALUES (?)", name)
			if err != nil {
				return fmt.Errorf("failed to insert category: %w", err)
			}
	
		case "package":
			if len(args) < 5 {
				return fmt.Errorf("%w:引数が足りません。例: bm add package <パッケージ名> <カテゴリID>", usageError)
			}
			name := args[3]
			var categoryID int
			categoryID, err = strconv.Atoi(args[4])
			if err != nil {
				return fmt.Errorf("%w:カテゴリIDは数値で指定してください。", usageError)
			}
			_, err = db.Exec("INSERT INTO packages (name, category_id) VALUES (?, ?)", name, categoryID)
			if err != nil {
				return fmt.Errorf("failed to insert package: %w", err)
			}
	
		default:
			return fmt.Errorf("%w:不明なサブコマンドです:", usageError)

		}
	
	default:
		return fmt.Errorf("%w:不明なコマンドです:", usageError)
	}
	return nil
}
