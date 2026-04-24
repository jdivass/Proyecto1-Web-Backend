package database

import (
	"database/sql"
)

func SeedDatabase(db *sql.DB)(err error){

	query := `select exists (select 1 from series limit 1)`
	var hasData bool
	 err = db.QueryRow(query).Scan(&hasData)
	 if err != nil {
		return err
	 }
	
	if !hasData {
		fillQuery := `
			insert into series (
				title,
				genre,
				description,
				platform,
				status,
				total_seasons,
				total_episodes,
				current_season,
				current_episode,
				image_path
			)
			values
			(
				'Breaking Bad',
				'Drama',
				'A chemistry teacher turns to producing methamphetamine after being diagnosed with cancer.',
				'Netflix',
				2,
				5,
				62,
				5,
				16,
				'uploads/breakingbad.jpg'
			),
			(
				'Dark',
				'Sci-Fi',
				'A mysterious disappearance reveals secrets across multiple generations in a small German town.',
				'Netflix',
				2,
				3,
				26,
				3,
				8,
				'uploads/dark.jpg'
			),
			(
				'Attack on Titan',
				'Anime',
				'Humanity fights for survival against giant humanoid creatures known as Titans.',
				'Crunchyroll',
				1,
				4,
				87,
				2,
				10,
				'uploads/aot.jpg'
			),
			(
				'Arcane',
				'Animation',
				'Two sisters become involved in conflict between the wealthy city of Piltover and the oppressed Zaun.',
				'Netflix',
				0,
				2,
				18,
				1,
				3,
				'uploads/arcane.jpg'
			),
			(
				'Stranger Things',
				'Fantasy',
				'A group of kids uncover supernatural mysteries in their small town.',
				'Netflix',
				1,
				4,
				34,
				2,
				5,
				'uploads/strangerthings.jpg'
			);
		`
		_, err = db.Exec(fillQuery)
		if err != nil {
			return err
		}	
	}
	return nil
}