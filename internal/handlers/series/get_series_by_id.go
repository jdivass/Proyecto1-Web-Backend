package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"backend/internal/models"
	"backend/internal/utils"
)

func GetSeriesById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			utils.WriteJSONError(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		idStr := r.PathValue("id")
		id, idErr := strconv.Atoi(idStr)
		if idErr != nil {
			utils.WriteJSONError(w, "invalid series id", http.StatusBadRequest)
			return
		}

		querySeries :=
			`SELECT id,
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
						WHERE id = ?
				`

		row := db.QueryRow(querySeries, id)

		var serie models.Series
		row_err := row.Scan(
			&serie.ID,
			&serie.Title,
			&serie.Genre,
			&serie.Description,
			&serie.Platform,
			&serie.Status,
			&serie.ImagePath,
			&serie.TotalSeasons,
			&serie.TotalEpisodes,
			&serie.CurrentSeason,
			&serie.CurrentEpisode,
			&serie.CreatedAt,
			&serie.UpdatedAt,
		)

		if row_err == sql.ErrNoRows {
			utils.WriteJSONError(w, "series not found", http.StatusNotFound)
			return
		}

		if row_err != nil {
			utils.WriteJSONError(w, "database query error", http.StatusInternalServerError)
			return
		}

		var rating models.Rating
		queryRating := `
				select id, series_id, content, stars_quantity, created_at
				from ratings
				where series_id = ?
			`
		ratingErr := db.QueryRow(queryRating, id).Scan(&rating.ID,
			&rating.SeriesID,
			&rating.Content,
			&rating.StarsQuantity,
			&rating.CreatedAt)

		if ratingErr != nil && ratingErr != sql.ErrNoRows {
			utils.WriteJSONError(w, "rating query error", http.StatusInternalServerError)
			return
		}

		if ratingErr == nil {
			serie.Rating = &rating
		}
		utils.WriteJSONResponse(w, serie, http.StatusOK)
	}
}
