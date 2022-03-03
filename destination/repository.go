package destination

import (
	"apitourism/bookmark"
	"apitourism/image"
	"apitourism/rating"
	"apitourism/view"
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Add(destination Destination) (Destination, error)
	DeleteByUUID(id string) error
	FindByName(name string) (Destination, error)
	GetByUUID(id string) (Destination, error)
	GetAllByLimitPage(limit int, page int) ([]Destination, error)
	FindNearby(inputGeohash string, inputType string, inputRating bool, inputPopularity bool) ([]NearbyDestination, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (this *repository) Add(destination Destination) (Destination, error) {
	// Check apa data sudah ada
	_, err := this.FindByName(destination.Name)
	if err == nil {
		return destination, errors.New("Destination already exist")
	}

	err = this.db.Create(&destination).Error
	if err != nil {
		return destination, err
	}

	return destination, nil
}

func (this *repository) DeleteByUUID(id string) error {
	// Cek apakah data yang ingin di delete ada
	_, err := this.GetByUUID(id)
	if err != nil {
		return err
	}

	// Lakukan delete
	err = this.db.Where("id = ?", id).Delete(&Destination{}).Error
	if err != nil {
		return err
	}

	///// Delete soft data di semua tabel yang memiliki relasi /////
	// Delete Bookmark
	err = this.db.Where("destination_id = ?", id).Delete(&bookmark.Bookmark{}).Error
	if err != nil {
		return err
	}
	// Delete Rating
	err = this.db.Where("destination_id = ?", id).Delete(&rating.Rating{}).Error
	if err != nil {
		return err
	}
	// Delete View
	err = this.db.Where("destination_id = ?", id).Delete(&view.View{}).Error
	if err != nil {
		return err
	}
	// Delete Image
	err = this.db.Where("destination_id = ?", id).Delete(&image.Image{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (this *repository) FindByName(name string) (Destination, error) {
	destination := Destination{}

	err := this.db.Where("name = ?", name).First(&destination).Error
	if err != nil {
		return destination, err
	}

	return destination, nil
}

func (this *repository) GetByUUID(id string) (Destination, error) {
	destination := Destination{}

	err := this.db.First(&destination, "id = ?", id).Error
	if err != nil {
		return destination, errors.New("Destination not found")
	}

	return destination, nil
}

func (this *repository) GetAllByLimitPage(limit int, page int) ([]Destination, error) {
	// Mengambil data destination berdasarkan limit beserta page
	destinations := []Destination{}
	err := this.db.Limit(limit).Offset(limit * page).Find(&destinations).Error
	if err != nil {
		return destinations, err
	}
	return destinations, nil
}

func (this *repository) FindNearby(inputGeohash string, inputType string, inputRating bool, inputPopularity bool) ([]NearbyDestination, error) {
	allDestinations := []NearbyDestination{}
	inputGeohash = inputGeohash[:len(inputGeohash)-1]

	subQuery1 := this.db.
		Select("COUNT(views.id) AS view_amount, destinations.id, destinations.type, destinations.geo_hash, destinations.name").
		Table("views").
		Joins("JOIN destinations ON destinations.id = views.destination_id").
		Group("destinations.id")
	subQuery2 := this.db.
		Select("AVG(ratings.rate) AS rating, destinations.id").
		Table("ratings").
		Joins("JOIN destinations ON destinations.id = ratings.destination_id").
		Group("destinations.id")
	for len(allDestinations) < 5 && len(inputGeohash) > 1 {
		inputGeohash = inputGeohash[:len(inputGeohash)-1]
		if inputType == "all" {
			if inputRating == false && inputPopularity == false {
				this.db.Table("(?) AS m1", subQuery1).
					Joins("JOIN (?) AS m2 ON m1.id = m2.id", subQuery2).
					Select("m1.view_amount, m2.rating, m1.id, m1.type, m1.geo_hash, m1.name").
					Where("geo_hash LIKE ?", "qw%").
					Find(&allDestinations)
			} else if inputRating == false && inputPopularity == true {
				this.db.Table("(?) AS m1", subQuery1).
					Joins("JOIN (?) AS m2 ON m1.id = m2.id", subQuery2).
					Select("m1.view_amount, m2.rating, m1.id, m1.type, m1.geo_hash, m1.name").
					Where("geo_hash LIKE ?", "qw%").
					Order("view_amount DESC").
					Find(&allDestinations)
			} else if inputRating == true && inputPopularity == false {
				this.db.Table("(?) AS m1", subQuery1).
					Joins("JOIN (?) AS m2 ON m1.id = m2.id", subQuery2).
					Select("m1.view_amount, m2.rating, m1.id, m1.type, m1.geo_hash, m1.name").
					Where("geo_hash LIKE ?", "qw%").
					Order("rating DESC").
					Find(&allDestinations)
			} else if inputRating == true && inputPopularity == true {
				this.db.Table("(?) AS m1", subQuery1).
					Joins("JOIN (?) AS m2 ON m1.id = m2.id", subQuery2).
					Select("m1.view_amount, m2.rating, m1.id, m1.type, m1.geo_hash, m1.name").
					Where("geo_hash LIKE ?", "qw%").
					Order("view_amount DESC, rating DESC").
					Find(&allDestinations)
			}
		} else {
			if inputRating == false && inputPopularity == false {
				this.db.Table("(?) AS m1", subQuery1).
					Joins("JOIN (?) AS m2 ON m1.id = m2.id", subQuery2).
					Select("m1.view_amount, m2.rating, m1.id, m1.type, m1.geo_hash, m1.name").
					Where("type = ? AND geo_hash LIKE ?", inputType, inputGeohash+"%").
					Find(&allDestinations)
			} else if inputRating == false && inputPopularity == true {
				this.db.Table("(?) AS m1", subQuery1).
					Joins("JOIN (?) AS m2 ON m1.id = m2.id", subQuery2).
					Select("m1.view_amount, m2.rating, m1.id, m1.type, m1.geo_hash, m1.name").
					Where("type = ? AND geo_hash LIKE ?", inputType, inputGeohash+"%").
					Order("view_amount DESC").
					Find(&allDestinations)
			} else if inputRating == true && inputPopularity == false {
				this.db.Table("(?) AS m1", subQuery1).
					Joins("JOIN (?) AS m2 ON m1.id = m2.id", subQuery2).
					Select("m1.view_amount, m2.rating, m1.id, m1.type, m1.geo_hash, m1.name").
					Where("type = ? AND geo_hash LIKE ?", inputType, inputGeohash+"%").
					Order("rating DESC").
					Find(&allDestinations)
			} else if inputRating == true && inputPopularity == true {
				this.db.Table("(?) AS m1", subQuery1).
					Joins("JOIN (?) AS m2 ON m1.id = m2.id", subQuery2).
					Select("m1.view_amount, m2.rating, m1.id, m1.type, m1.geo_hash, m1.name").
					Where("type = ? AND geo_hash LIKE ?", inputType, inputGeohash+"%").
					Order("rating DESC, view_amount DESC").
					Find(&allDestinations)
			}
		}
	}

	return allDestinations, nil
}

// type all, popularity false, rating false
// type all, popularity true, rating false
// type all, popularity false, rating true
// type all, popularity true, rating true

// type ada, popularity false, rating false
// type ada, popularity true, rating false
// type ada, popularity false, rating true
// type ada, popularity true, rating true
