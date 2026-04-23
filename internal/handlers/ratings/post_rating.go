package handlers

import (
	"net/http"
	"database/sql"
	"backend/internal/models"
	"backend/internal/utils"
	"encoding/json"

)

func CreateRating(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			utils.WriteJSONError(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var rating models.Rating

		decodeErr := json.NewDecoder(r.Body).Decode(&rating)
		if decodeErr != nil {
			utils.WriteJSONError(w, "invalid request body", http.StatusBadRequest)
			return
		}
		validateErr := utils.ValidateRatings(rating)

		if validateErr != nil {
			utils.WriteJSONError(w, validateErr.Error(), http.StatusBadRequest)
			return
		}

		var exists int
		seriesErr := db.QueryRow("select id from series WHERE id = ?", rating.SeriesID).Scan(&exists)

		if seriesErr == sql.ErrNoRows {
			utils.WriteJSONError(w, "series not found", http.StatusNotFound)
			return
		}

		if seriesErr != nil {
			utils.WriteJSONError(w, "database error", http.StatusInternalServerError)
			return
		}

		var existing int
		existingErr := db.QueryRow(`select id from ratings where series_id = ?`, rating.SeriesID).Scan(&existing)

		if existingErr == nil {
			utils.WriteJSONError(w,"rating already exists", http.StatusConflict)
			return
		}

		if existingErr != sql.ErrNoRows {
			utils.WriteJSONError(w, "database query error", http.StatusInternalServerError)
			return
		}

		insertQuery := `
			insert into ratings (series_id, content, stars_quantity)
			values (?, ?, ?)
		`

		insertResult, insertErr := db.Exec(insertQuery,
			rating.SeriesID,
			rating.Content,
			rating.StarsQuantity,
		)

		if insertErr != nil {
			utils.WriteJSONError(w, "could not create rating", http.StatusInternalServerError)
			return
		}

		id, _ := insertResult.LastInsertId()

		rating.ID = int(id)

		utils.WriteJSONResponse(w, rating, http.StatusCreated)
	}
}