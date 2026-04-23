package utils

import (
	"errors"
	"backend/internal/models"
)

func ValidateRatings(rating models.Rating) error {
	if rating.StarsQuantity < 1 || rating.StarsQuantity > 5 {
		return errors.New("stars must be a value between 1 and 5")
	}

	if rating.SeriesID <= 0 {
		return errors.New("series_id is required")
	}

	if len(rating.Content) > 500 {
		return errors.New("content is too long")
	}

	return nil
}