package user

import (
	"apitourism/bookmark"
	"apitourism/rating"
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
	DeleteByUUID(id string) error
	FindByEmail(email string) (User, error)
	GetByUUID(id string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (this *repository) Save(user User) (User, error) {
	_, err := this.FindByEmail(user.Email)
	if err == nil {
		return user, errors.New("email has been used")
	}

	err = this.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (this *repository) DeleteByUUID(id string) error {
	// Cek apakah data yang ingin di delete ada
	_, err := this.GetByUUID(id)
	if err != nil {
		return err
	}

	// Lakukan delete
	err = this.db.Where("id = ?", id).Delete(&User{}).Error
	if err != nil {
		return err
	}

	///// Delete soft data di semua tabel yang memiliki relasi /////
	// Delete Bookmark
	err = this.db.Where("user_id = ?", id).Delete(&bookmark.Bookmark{}).Error
	if err != nil {
		return err
	}
	// Delete Rating
	err = this.db.Where("user_id = ?", id).Delete(&rating.Rating{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (this *repository) FindByEmail(email string) (User, error) {
	user := User{}

	err := this.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, errors.New("Email has not been registered")
	}

	return user, nil
}

func (this *repository) GetByUUID(id string) (User, error) {
	user := User{}

	err := this.db.First(&user, "id = ?", id).Error
	if err != nil {
		return user, errors.New("User not found")
	}

	return user, nil
}

// userHandler.RegisterUserHandler()
// service.RegisterUser()	<-- service adalah atribut dari userHandler
// repository.Save()	<-- repository adalah atribut dari service
// db.Create()	<-- db adalah atribut dari repository

/// CARA INSTANSIASI
// buat objek dari repository dengan mempassing objek dari db yg sudah dibuat
// buat objek dari service dengan mempassing objek dari repository yg sudah dibuat
// buat objek dari userHandler dengan mempassing objek dari service yg sudah dibuat
