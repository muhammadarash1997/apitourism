package rating

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Add(rating Rating) (Rating, error)
	FindByDestinationIDAndUserID(destinationID string, userID string) (Rating, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (this *repository) Add(rating Rating) (Rating, error) {
	// Check DestinationID And UserID
	_, err := this.FindByDestinationIDAndUserID(rating.DestinationID, rating.UserID)

	if err == nil {
		return rating, errors.New("Data already exist")
	}

	err = this.db.Create(&rating).Error
	if err != nil {
		return rating, err
	}

	return rating, nil
}

func (this *repository) FindByDestinationIDAndUserID(destinationID string, userID string) (Rating, error) {
	rating := Rating{}

	err := this.db.Where("destination_id = ? AND user_id = ?", destinationID, userID).First(&rating).Error
	if err != nil {
		return rating, err
	}

	return rating, nil
}
