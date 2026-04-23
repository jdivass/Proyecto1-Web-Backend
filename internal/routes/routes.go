package routes

import (
	"net/http"
	"database/sql"
	"backend/internal/handlers/ratings"
	"backend/internal/handlers/series"
)
func RegisterRoutes(mux *http.ServeMux, db*sql.DB){
	mux.HandleFunc("GET /series", series.GetSeries(db))
	mux.HandleFunc("GET /series/{id}", series.GetSeriesById(db))
	mux.HandleFunc("POST /series", series.CreateSeries(db))
	mux.HandleFunc("DELETE /series/{id}", series.DeleteSeries(db))
	mux.HandleFunc("PUT /series/{id}", series.UpdateSeries(db))
	mux.HandleFunc("POST /series/{id}/rating", ratings.CreateRating(db))
	mux.HandleFunc("PUT /series/{id}/rating", ratings.UpdateRating(db))
	mux.HandleFunc("DELETE /series/{id}/rating", ratings.DeleteRating(db))
}