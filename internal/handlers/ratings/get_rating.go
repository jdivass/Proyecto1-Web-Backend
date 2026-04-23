package handlers

import (
	"database/sql"
	"net/http"
)
func GetRating(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodGet {

		}
	}
}