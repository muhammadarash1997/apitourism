package bookmark

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Add(bookmark Bookmark) (Bookmark, error)
	DeleteByUUID(id string) error
	FindByUserIDAndDestinationID(userID string, destinationID string) (Bookmark, error)
	GetByUUID(id string) (Bookmark, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (this *repository) Add(bookmark Bookmark) (Bookmark, error) {
	// Check UserID
	_, err := this.FindByUserIDAndDestinationID(bookmark.UserID, bookmark.DestinationID)

	if err == nil {
		return bookmark, errors.New("Data already exist")
	}

	err = this.db.Create(&bookmark).Error
	if err != nil {
		return bookmark, err
	}

	return bookmark, nil
}

func (this *repository) DeleteByUUID(id string) error {
	// Cek apakah data yang ingin di delete ada
	_, err := this.GetByUUID(id)
	if err != nil {
		return err
	}

	// Lakukan delete
	err = this.db.Where("id = ?", id).Delete(&Bookmark{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (this *repository) FindByUserIDAndDestinationID(userID string, destinationID string) (Bookmark, error) {
	bookmark := Bookmark{}

	err := this.db.Where("user_id = ? AND destination_id = ?", userID, destinationID).First(&bookmark).Error
	if err != nil {
		return bookmark, err
	}

	return bookmark, nil
}

func (this *repository) GetByUUID(id string) (Bookmark, error) {
	bookmark := Bookmark{}

	err := this.db.First(&bookmark, "id = ?", id).Error
	if err != nil {
		return bookmark, errors.New("Bookmark not found")
	}

	return bookmark, nil
}