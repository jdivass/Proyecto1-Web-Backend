package models
import (
	"time"
)

type Series struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Genre string `json:"genre"`
	Description string `json:"description"`
	Platform string `json:"platform"`
	Status int `json:"status"`
	ImagePath string `json:"image_path"`

	TotalSeasons int `json:"total_seasons"`
	TotalEpisodes int `json:"total_episodes"`

	CurrentSeason int `json:"current_season"`
	CurrentEpisode int `json:"current_episode"`
 
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Rating *Rating `json:"rating,omitempty"`
}