package series

import (
	"database/sql"
	"math"
	"net/http"
	"strconv"

	"backend/internal/models"
	"backend/internal/utils"
)

func GetSeries(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			utils.WriteJSONError(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		pageStr := r.URL.Query().Get("page")

		if pageStr == "" {
			pageStr = "1"
		}

		limitStr := r.URL.Query().Get("limit")

		if limitStr == "" {
			limitStr = "10"
		}

		pageInt, pageErr := strconv.Atoi(pageStr)

		if pageErr != nil {
			utils.WriteJSONError(w, "error getting page", http.StatusInternalServerError)
			return
		}

		limitInt, limitErr := strconv.Atoi(limitStr)

		if limitErr != nil {
			utils.WriteJSONError(w, "error getting limit", http.StatusInternalServerError)
			return
		}

		offset := (pageInt - 1) * limitInt

		queryTotal := `
			select count(*)
			from series
			`
		var totalCount int
		totalErr := db.QueryRow(queryTotal).Scan(&totalCount)

		if totalErr != nil {
			utils.WriteJSONError(w, "error getting total series", http.StatusInternalServerError)
			return
		}

		query := `
			select
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
			from series
			order by id desc
			limit ?
			offset ?
		`
		rows, err := db.Query(query, limitInt, offset)
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

		total_pages := math.Ceil(float64(totalCount) / float64(limitInt))
		utils.WriteJSONResponse(w, map[string]any{
			"data":        seriesList,
			"total":       totalCount,
			"page":        pageInt,
			"limit":       limitInt,
			"total_pages": total_pages,
		}, http.StatusOK)
	}
}
