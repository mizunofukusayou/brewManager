package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func main() {
	// データベースファイルを開く（存在しない場合は作成される）
	db, err := sql.Open("sqlite", "brewmanager.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 接続確認
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("SQLiteデータベースへの接続に成功しました")
}
