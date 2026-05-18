package main

import (
	"encoding/json"
	"net/http"
)

func (a *api) getPackages(w http.ResponseWriter, r *http.Request) {

	rows, err := a.db.Query("SELECT packages.ID, categories.name AS category_name, packages.name, packages.notes FROM packages JOIN categories ON packages.category_id = categories.id")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{"データの取得に失敗しました"})
		return
	}
	defer rows.Close()
	type row struct {
		ID           int    `json:"id"`
		CategoryName string `json:"category_name"`
		Name         string `json:"name"`
		Notes        string `json:"notes"`
	}

	type response struct {
		Packages []row `json:"packages"`
	}

	var res response
	for rows.Next() {
		var row row
		err := rows.Scan(&row.ID, &row.CategoryName, &row.Name, &row.Notes)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{"データの読み込みに失敗しました"})
			return
		}
		res.Packages = append(res.Packages, row)
	}
	if err := rows.Err(); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{"データ取得中にエラーが発生しました"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
