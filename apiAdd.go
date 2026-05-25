package main

import (
	"encoding/json"
	"net/http"
)

func (a *api) apiAdd(w http.ResponseWriter, r *http.Request) {
	table := r.FormValue("table")
	switch table {
	case "package":
		name := r.FormValue("name")
		categoryID := r.FormValue("category_id")
		_, err := a.db.Exec("INSERT INTO packages (name, category_id) VALUES (?, ?)", name, categoryID)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{"データの追加に失敗しました"})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Package added successfully")
	case "category":
		name := r.FormValue("name")
		_, err := a.db.Exec("INSERT INTO categories (name) VALUES (?)", name)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{"データの追加に失敗しました"})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Category added successfully")
	default:
		http.Error(w, "Invalid table specified", http.StatusBadRequest)
	}
}
