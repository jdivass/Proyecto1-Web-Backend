package ratings

import (
	"database/sql"
	"net/http"
	"strconv"

	"backend/internal/utils"
)

func DeleteRating(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodDelete {
			utils.WriteJSONError(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		idStr := r.PathValue("id")

		id, convErr := strconv.Atoi(idStr)
		if convErr != nil {
			utils.WriteJSONError(w, "invalid series id", http.StatusBadRequest)
			return
		}

		result, deleteErr := db.Exec(
			"delete from ratings where series_id = ?",
			id,
		)

		if deleteErr != nil {
			utils.WriteJSONError(w, "database error", http.StatusInternalServerError)
			return
		}

		rowsAffected, _ := result.RowsAffected()

		if rowsAffected == 0 {
			utils.WriteJSONError(w, "rating not found", http.StatusNotFound)
			return
		}

		utils.WriteJSONResponse(w, map[string]string{
			"message": "rating deleted successfully",
		}, http.StatusOK)
	}
}