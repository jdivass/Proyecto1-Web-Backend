package database
import (
	"database/sql"
)

func InitializeDB(db *sql.DB) (err error) {
	query := `
	PRAGMA foreign_keys = ON;
	create table if not exists series (
		id integer primary key autoincrement,
		title text not null,
		genre text not null,
		description text not null,
		platform text not null,
		status integer not null default 0 check (status in (0,1,2)),

		total_seasons int not null,
		total_episodes int not null,
		
		current_season integer not null,
		current_episode integer not null,
		
		created_at datetime default current_timestamp,
		updated_at datetime default current_timestamp,

		image_path text
	);

	create table if not exists ratings (
		id integer primary key autoincrement,
		series_id integer unique references series(id) on delete cascade,
		content text,
		stars_quantity integer not null,
		constraint check_stars check (stars_quantity between 1 and 5),
		created_at datetime default current_timestamp
	);
	`
	_, err = db.Exec(query)

	return err
}