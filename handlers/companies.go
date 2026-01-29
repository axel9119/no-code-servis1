package handlers

import (
	"encoding/json"
	"net/http"
)

// Пример структуры Company (можно заменить на модель с Postgres позже)
type Company struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	Components string `json:"components"`
}

// Заглушка для GET /api/companies
func GetCompanies(w http.ResponseWriter, r *http.Request) {
	companies := []Company{
		{ID: "1", Name: "My Company", Slug: "my-company", Components: "{}"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(companies)
}
