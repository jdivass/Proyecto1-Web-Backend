package ratings

import (
	"net/http"
	"database/sql"
	"backend/internal/models"
	"backend/internal/utils"
	"encoding/json"
	"strconv"
)

func UpdateRating(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPut {
			utils.WriteJSONError(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		idStr := r.PathValue("id")

		id, convErr := strconv.Atoi(idStr)
		if convErr != nil {
			utils.WriteJSONError(w, "invalid rating id reference", http.StatusBadRequest)
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

		rating.SeriesID = id

		var existingID int

		scanErr := db.QueryRow("select id from ratings where series_id = ?", id).Scan(&existingID)

		if scanErr == sql.ErrNoRows {
			query := `
				insert into ratings (series_id, content, stars_quantity)
				values (?, ?, ?)
			`

			result, insertErr := db.Exec(
				query,
				rating.SeriesID,
				rating.Content,
				rating.StarsQuantity,
			)

			if insertErr != nil {
				utils.WriteJSONError(w, "could not create rating", http.StatusInternalServerError)
				return
			}

			newID, _ := result.LastInsertId()
			rating.ID = int(newID)

			utils.WriteJSONResponse(w, rating, http.StatusCreated)
			return
		}

		if scanErr != nil {
			utils.WriteJSONError(w, "database error", http.StatusInternalServerError)
			return
		}

		updateQuery := `
			update ratings
			set content = ?,
				stars_quantity = ?
			where series_id = ?
		`

		_, updateErr := db.Exec(
			updateQuery,
			rating.Content,
			rating.StarsQuantity,
			id,
		)

		if updateErr != nil {
			utils.WriteJSONError(w, "could not update rating", http.StatusInternalServerError)
			return
		}

		rating.ID = existingID

		utils.WriteJSONResponse(w, rating, http.StatusOK)
	}
}