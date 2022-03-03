package image

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Save(image Image) (Image, error)
	DeleteByUUID(id string) error
	GetByUUID(id string) (Image, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (this *repository) Save(image Image) (Image, error) {
	err := this.db.Create(&image).Error
	if err != nil {
		return image, err
	}

	return image, nil
}

func (this *repository) DeleteByUUID(id string) error {
	// Cek apakah data yang ingin di delete ada
	_, err := this.GetByUUID(id)
	if err != nil {
		return err
	}

	// Lakukan delete
	err = this.db.Where("id = ?", id).Delete(&Image{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (this *repository) GetByUUID(id string) (Image, error) {
	image := Image{}

	err := this.db.First(&image, "id = ?", id).Error
	if err != nil {
		return image, errors.New("Image not found")
	}

	return image, nil
}
