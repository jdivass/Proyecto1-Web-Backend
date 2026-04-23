package handlers

import (
	"backend/internal/models"
	"backend/internal/utils"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
)

func UpdateSeries(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			utils.WriteJSONError(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		defer r.Body.Close()

		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			utils.WriteJSONError(w, "invalid series id", http.StatusBadRequest)
			return
		}

		var series models.Series
		err = json.NewDecoder(r.Body).Decode(&series)

		if err != nil {
			utils.WriteJSONError(w, "invalid request body", http.StatusBadRequest)
			return
		}

		seriesErr := utils.ValidateSeries(series)

		if seriesErr != nil {
			utils.WriteJSONError(w, seriesErr.Error(), http.StatusBadRequest)
			return
		}

		row := db.QueryRow("select id from series where id = ?", id)

		var existingID int
		err = row.Scan(&existingID)

		if err == sql.ErrNoRows {
			utils.WriteJSONError(w, "cannot find that series", http.StatusNotFound)
			return
		}

		if err != nil {
			utils.WriteJSONError(w, "database query error", http.StatusInternalServerError)
			return
		}

		updateQuery := `
			update series
			set
				title = ?,
				genre = ?,
				description = ?,
				platform = ?,
				status = ?,
				image_path = ?,
				total_seasons = ?,
				total_episodes = ?,
				current_season = ?,
				current_episode = ?,
				updated_at = current_timestamp
			where id = ?
		`
		updateResult, updateQueryErr := db.Exec(updateQuery,
			series.Title,
			series.Genre,
			series.Description,
			series.Platform,
			series.Status,
			series.ImagePath,
			series.TotalSeasons,
			series.TotalEpisodes,
			series.CurrentSeason,
			series.CurrentEpisode, id)

		if updateQueryErr != nil {
			utils.WriteJSONError(w, "database query error", http.StatusInternalServerError)
			return
		}

		rowsAffected, _ := updateResult.RowsAffected()

		if rowsAffected == 0 {
			utils.WriteJSONError(w, "cannot find that series", http.StatusNotFound)
			return
		}

		utils.WriteJSONResponse(w, "series updated successfully", http.StatusOK)
	}
}
