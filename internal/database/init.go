package database

import (
	"database/sql"
	"log"
)

func InitializeDB(db *sql.DB) (err error) {
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		log.Println("Pragma ERROR", err)
		return err
	}
	querySeries := `
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
	)`
	_, err = db.Exec(querySeries)
	if err != nil {
		log.Println("series table ERROR", err)
		return err
	}
	queryRatings := `
		create table if not exists ratings (
		id integer primary key autoincrement,
		series_id integer unique references series(id) on delete cascade,
		content text,
		stars_quantity integer not null,
		created_at datetime default current_timestamp,
		constraint check_stars check (stars_quantity between 1 and 5)
	)
	`
	_, err = db.Exec(queryRatings)
	if err != nil {
		log.Println("ratings table ERROR", err)
		return err
	}

	return nil
}
