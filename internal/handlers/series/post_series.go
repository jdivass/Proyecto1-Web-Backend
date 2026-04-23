package handlers

import (
	"backend/internal/models"
	"backend/internal/utils"
	"database/sql"
	"encoding/json"
	"net/http"
)

func CreateSeries(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodPost {
			utils.WriteJSONError(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		defer r.Body.Close()
		var series models.Series
		err := json.NewDecoder(r.Body).Decode(&series)
		
		if err != nil {
			utils.WriteJSONError(w, "invalid request body", http.StatusBadRequest)
			return
		}
		seriesErr := utils.ValidateSeries(series)

		if seriesErr != nil {
			utils.WriteJSONError(w, seriesErr.Error(), http.StatusBadRequest)
			return
		}

		query := `
			insert into series (
				title,
				genre,
				description,
				platform,
				status,
				image_path,
				total_seasons,
				total_episodes,
				current_season,
				current_episode	
			)
			
			values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`

		result, insertErr := db.Exec(query,
							series.Title,
							series.Genre,
							series.Description,
							series.Platform,
							series.Status,
							series.ImagePath,
							series.TotalSeasons,
							series.TotalEpisodes,
							series.CurrentSeason,
							series.CurrentEpisode)
		
		if insertErr != nil{
			utils.WriteJSONError(w, "database query error", http.StatusInternalServerError)
			return
		}

		id, lastInsertErr := result.LastInsertId()
		
		if lastInsertErr != nil {
			utils.WriteJSONError(w, "database query error", http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{
			"id": id,
			"message": "series created successfully",
		}
		utils.WriteJSONResponse(w, response, http.StatusCreated)
	}
}	