package main

import "fmt"

func handleList(args []string) error {
	if len(args) < 3 {
		return fmt.Errorf("%w:引数が足りません。例: bm list <テーブル名>", usageError)
	}

	db, err := getDB()
	if err != nil {
		return fmt.Errorf("failed to get database: %w", err)
	}
	defer db.Close()

	tableName := args[2]

	switch tableName {
	case "categories":
		rows, err := db.Query("SELECT * FROM categories")
		if err != nil {
			return fmt.Errorf("%sからの取得に失敗しました: %w", tableName, err)
		}
		defer rows.Close()
		for rows.Next() {
			var row Category
			err := rows.Scan(&row.ID, &row.Name)
			if err != nil {
				return fmt.Errorf("failed to scan category: %w", err)
			}
			fmt.Printf("ID: %d, Name: %s\n", row.ID, row.Name)
		}
		if err := rows.Err(); err != nil {
			return fmt.Errorf("データ取得中にエラーが発生しました: %w", err)
		}

	case "packages":
		rows, err := db.Query("SELECT packages.*, categories.name AS category_name FROM packages JOIN categories ON packages.category_id = categories.id")
		if err != nil {
			return fmt.Errorf("%sからの取得に失敗しました: %w", tableName, err)
		}
		defer rows.Close()
		for rows.Next() {
			type Row struct {
				Package
				CategoryName string
			}
			var row Row
			err := rows.Scan(&row.ID, &row.Name, &row.CategoryID, &row.Notes, &row.CategoryName)
			if err != nil {
				return fmt.Errorf("failed to scan package: %w", err)
			}
			if row.Notes == "" {
				fmt.Printf("ID: %d, Category: %s, Name: %s\n", row.ID, row.CategoryName, row.Name)
			} else {
				fmt.Printf("ID: %d, Category: %s, Name: %s, Notes: %s\n", row.ID, row.CategoryName, row.Name, row.Notes)
			}
		}
		if err := rows.Err(); err != nil {
			return fmt.Errorf("データ取得中にエラーが発生しました: %w", err)
		}

	default:
		return fmt.Errorf("%w: unknown table name: %s", usageError, tableName)
	}

	return nil
}
