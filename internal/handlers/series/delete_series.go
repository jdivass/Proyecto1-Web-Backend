package handlers

import (
	"backend/internal/utils"
	"database/sql"
	"net/http"
	"strconv"
)
func DeleteSeries(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodDelete {
			utils.WriteJSONError(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			utils.WriteJSONError(w, "invalid series id", http.StatusBadRequest)
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

		deleteResult, deleteErr := db.Exec("delete from series where id = ?", id)

		if deleteErr != nil {
			utils.WriteJSONError(w, "database query error", http.StatusInternalServerError)
			return
		}

		rowsAffected, _ := deleteResult.RowsAffected()

		if rowsAffected == 0 {
			utils.WriteJSONError(w, "series not found", http.StatusNotFound)
			return
		}

		utils.WriteJSONResponse(w, "series deleted successfully", http.StatusOK)
	}
}