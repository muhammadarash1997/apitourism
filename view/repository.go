package view

import "gorm.io/gorm"

type Repository interface {
	Add(view View) (View, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (this *repository) Add(view View) (View, error) {
	err := this.db.Create(&view).Error
	if err != nil {
		return view, err
	}

	return view, nil
}