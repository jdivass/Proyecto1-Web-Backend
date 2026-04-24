package series

import (
	"backend/internal/models"
	"backend/internal/utils"
	"database/sql"
	"io"
	"net/http"
	"strconv"
	"encoding/base64"
)

func UpdateSeries(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			utils.WriteJSONError(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			utils.WriteJSONError(w, "invalid series id", http.StatusBadRequest)
			return
		}

		err = r.ParseMultipartForm(10 << 20)
		if err != nil {
			utils.WriteJSONError(w, "invalid multipart form", http.StatusBadRequest)
			return
		}

		var series models.Series
		series.Title = r.FormValue("title")
		series.Genre = r.FormValue("genre")
		series.Description = r.FormValue("description")
		series.Platform = r.FormValue("platform")

		status, err := strconv.Atoi(r.FormValue("status"))
		if err != nil {
			utils.WriteJSONError(w, "invalid status", http.StatusBadRequest)
			return
		}
		series.Status = status

		totalSeasons, err := strconv.Atoi(r.FormValue("total_seasons"))
		if err != nil {
			utils.WriteJSONError(w, "invalid total_seasons", http.StatusBadRequest)
			return
		}
		series.TotalSeasons = totalSeasons

		totalEpisodes, err := strconv.Atoi(r.FormValue("total_episodes"))
		if err != nil {
			utils.WriteJSONError(w, "invalid total_episodes", http.StatusBadRequest)
			return
		}
		series.TotalEpisodes = totalEpisodes

		currentSeason, err := strconv.Atoi(r.FormValue("current_season"))
		if err != nil {
			utils.WriteJSONError(w, "invalid current_season", http.StatusBadRequest)
			return
		}
		series.CurrentSeason = currentSeason

		currentEpisode, err := strconv.Atoi(r.FormValue("current_episode"))
		if err != nil {
			utils.WriteJSONError(w, "invalid current_episode", http.StatusBadRequest)
			return
		}
		series.CurrentEpisode = currentEpisode

		var currentImagePath string
		db.QueryRow("select image_path from series where id = ?", id).Scan(&currentImagePath)

		imagePath := currentImagePath

		file, _, fileErr := r.FormFile("image")
		if fileErr == nil {
			defer file.Close()
			imageBytes, readErr := io.ReadAll(file)
			if readErr == nil {
				mime := http.DetectContentType(imageBytes)
				b64 := base64.StdEncoding.EncodeToString(imageBytes)
				imagePath = "data:" + mime + ";base64," + b64
			}
		}

		series.ImagePath = imagePath

		validateErr := utils.ValidateSeries(series)
		if validateErr != nil {
			utils.WriteJSONError(w, validateErr.Error(), http.StatusBadRequest)
			return
		}

		var existingID int
		scanErr := db.QueryRow("select id from series where id = ?", id).Scan(&existingID)
		if scanErr == sql.ErrNoRows {
			utils.WriteJSONError(w, "series not found", http.StatusNotFound)
			return
		}
		if scanErr != nil {
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
		_, updateErr := db.Exec(updateQuery,
			series.Title,
			series.Genre,
			series.Description,
			series.Platform,
			series.Status,
			series.ImagePath,
			series.TotalSeasons,
			series.TotalEpisodes,
			series.CurrentSeason,
			series.CurrentEpisode,
			id,
		)

		if updateErr != nil {
			utils.WriteJSONError(w, "database query error", http.StatusInternalServerError)
			return
		}

		utils.WriteJSONResponse(w, "series updated successfully", http.StatusOK)
	}
}