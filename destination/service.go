package destination

import (
	"apitourism/helper"
)

type Service interface {
	AddDestination(destinationinput DestinationInput) (Destination, error)
	DeleteDestinationByUUID(id string) error
	GetAllDestinationsByLimitPage(limit int, page int) ([]Destination, error)
	FindNearbyDestination(inputCoordinate string, inputType string, inputRating bool, inputPopularity bool) ([]NearbyDestination, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (this *service) AddDestination(destinationInput DestinationInput) (Destination, error) {
	destination := Destination{}

	destination.Name = destinationInput.Name
	destination.Type = destinationInput.Type

	geohashEncoded := helper.EncodeGeohash(destinationInput.GeoHash)

	destination.GeoHash = geohashEncoded

	destinationAdded, err := this.repository.Add(destination)
	if err != nil {
		return destinationAdded, err
	}

	return destinationAdded, nil
}

func (this *service) DeleteDestinationByUUID(id string) error {
	err := this.repository.DeleteByUUID(id)
	if err != nil {
		return err
	}

	return nil
}

func (this *service) GetAllDestinationsByLimitPage(limit int, page int) ([]Destination, error) {
	destinations, err := this.repository.GetAllByLimitPage(limit, page)
	if err != nil {
		return destinations, err
	}

	return destinations, nil
}

func (this *service) FindNearbyDestination(inputCoordinate string, inputType string, inputRating bool, inputPopularity bool) ([]NearbyDestination, error) {
	geohash := helper.EncodeGeohash(inputCoordinate)
	// inputCoordinate = "qw67"
	destinations, err := this.repository.FindNearby(geohash, inputType, inputRating, inputPopularity)
	if err != nil {
		return destinations, err
	}

	return destinations, nil
}