package main

import (
	"fmt"
	"strconv"
)

func handleAdd(args []string) error {
	db, err := getDB()
	if err != nil {
		return fmt.Errorf("failed to get database: %w", err)
	}
	defer db.Close()
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
		fmt.Println("Added category:", name)

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
		fmt.Printf("Added package: {name: %s, category_id: %d}\n", name, categoryID)

	default:
		return fmt.Errorf("%w:%sは不明なサブコマンドです", usageError, args[2])
	}

	return nil
}
