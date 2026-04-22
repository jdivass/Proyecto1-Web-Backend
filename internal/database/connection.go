package database
import (
	"database/sql"
	_ "modernc.org/sqlite"
)
func ConnectDB() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite", "file:series.db")
	if err != nil {
		return nil, err
	}
	ping_err := db.Ping()
	if ping_err != nil {
		return nil, ping_err
	} 
	return db, nil
}

