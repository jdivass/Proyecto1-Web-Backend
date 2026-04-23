package handlers

import (
	"database/sql"
	"net/http"

	"backend/internal/models"
	"backend/internal/utils"
)

func GetSeries(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			utils.WriteJSONError(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		query := `
			SELECT
				id,
				title,
				genre,
				description,
				platform,
				status,
				image_path,
				total_seasons,
				total_episodes,
				current_season,
				current_episode,
				created_at,
				updated_at
			FROM series
			ORDER BY id DESC
		`

		rows, err := db.Query(query)
		if err != nil {
			utils.WriteJSONError(w, "database query error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var seriesList []models.Series

		for rows.Next() {
			var series models.Series

			err := rows.Scan(
				&series.ID,
				&series.Title,
				&series.Genre,
				&series.Description,
				&series.Platform,
				&series.Status,
				&series.ImagePath,
				&series.TotalSeasons,
				&series.TotalEpisodes,
				&series.CurrentSeason,
				&series.CurrentEpisode,
				&series.CreatedAt,
				&series.UpdatedAt,
			)

			if err != nil {
				utils.WriteJSONError(w, "error reading database rows", http.StatusInternalServerError)
				return
			}
			seriesList = append(seriesList, series)
		}

		if err = rows.Err(); err != nil {
			utils.WriteJSONError(w, "rows iteration error", http.StatusInternalServerError)
			return
		}

		utils.WriteJSONResponse(w, seriesList, http.StatusOK)
	}
}
