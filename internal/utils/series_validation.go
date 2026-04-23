package utils

import (
	"errors"

	"backend/internal/models"
)

func ValidateSeries(series models.Series) error {
	if series.Title == "" {
		return errors.New("title is required")
	}

	if series.Genre == "" {
		return errors.New("genre is required")
	}

	if series.Description == "" {
		return errors.New("description is required")
	}

	if series.Platform == "" {
		return errors.New("platform is required")
	}

	if series.ImagePath == "" {
		return errors.New("image path is required")
	}

	if series.Status < 0 || series.Status > 2 {
		return errors.New("status must be a value between 0 and 2")
	}

	if series.TotalSeasons <= 0 {
		return errors.New("total seasons must be a positive number")
	}

	if series.TotalEpisodes <= 0 {
		return errors.New("total episodes must be a positive number")
	}

	if series.CurrentSeason <= 0 {
		return errors.New("current season must be a positive number")
	}

	if series.CurrentEpisode <= 0 {
		return errors.New("current episode must be a positive number")
	}

	if series.CurrentSeason > series.TotalSeasons {
		return errors.New("current season cannot be greater than total seasons")
	}

	if series.CurrentEpisode > series.TotalEpisodes {
		return errors.New("current episode cannot be greater than total episodes")
	}

	return nil
}