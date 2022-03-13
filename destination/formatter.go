package destination

import (
	"apitourism/helper"
	"fmt"
)

type DestinationAddedFormatter struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	GeoHash string `json:"geo_hash"`
}

type DestinationGottenFormatter struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	GeoHash  string `json:"geo_hash"`
	Distance string `json:"distance"`
}

type NearbyDestinationGottenFormatter struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Type       string  `json:"type"`
	GeoHash    string  `json:"geo_hash"`
	ViewAmount int     `json:"view_amount"`
	Rating     float64 `json:"rating"`
	Distance   string  `json:"distance"`
}

func FormatDestinationAdded(destination Destination) DestinationAddedFormatter {
	destinationFormatted := DestinationAddedFormatter{
		ID:      destination.ID,
		Name:    destination.Name,
		Type:    destination.Type,
		GeoHash: destination.GeoHash,
	}

	return destinationFormatted
}

func FormatDestinationsGotten(userCoordinate string, allDestinationsGotten []Destination) []DestinationGottenFormatter {
	// Parse user coordinate
	userLatitude, userLongitude := helper.CoordinateParser(userCoordinate)

	allDestinationsFormatted := []DestinationGottenFormatter{}

	for _, destinationGotten := range allDestinationsGotten {
		// Create destinationFormatted
		destinationFormatted := DestinationGottenFormatter{}

		// Decode Geohash dari setiap objek dari destinations
		destinationLatitude, destinationLongitude := helper.DecodeGeohash(destinationGotten.GeoHash)

		// Calculate Distance dari setiap objek dari destinations
		distance := helper.CalculateDistance(userLatitude, userLongitude, destinationLatitude, destinationLongitude)
		distanceString := fmt.Sprintf("%g", distance) + " " + "km"

		// Masukkan hasil distance ke field Distancenya destinationFormatted
		destinationFormatted.ID = destinationGotten.ID
		destinationFormatted.Name = destinationGotten.Name
		destinationFormatted.Type = destinationGotten.Type
		destinationFormatted.GeoHash = destinationGotten.GeoHash
		destinationFormatted.Distance = distanceString

		// Tambahkan destinationFormatted ke allDestinationsFormatted
		allDestinationsFormatted = append(allDestinationsFormatted, destinationFormatted)
	}

	return allDestinationsFormatted
}

func FormatNearbyDestinationsGotten(userCoordinate string, allNearbyDestinationsGotten []NearbyDestination) []NearbyDestinationGottenFormatter {
	// Parse user coordinate
	userLatitude, userLongitude := helper.CoordinateParser(userCoordinate)

	allNearbyDestinationsFormatted := []NearbyDestinationGottenFormatter{}

	for _, nearbyDestinationGotten := range allNearbyDestinationsGotten {
		// Create destinationFormatted
		nearbyDestinationFormatted := NearbyDestinationGottenFormatter{}

		// Decode Geohash dari setiap objek dari destinations
		destinationLatitude, destinationLongitude := helper.DecodeGeohash(nearbyDestinationGotten.GeoHash)

		// Calculate Distance dari setiap objek dari destinations
		distance := helper.CalculateDistance(userLatitude, userLongitude, destinationLatitude, destinationLongitude)
		distanceString := fmt.Sprintf("%g", distance) + " " + "km"

		// Masukkan hasil distance ke field Distancenya destinationFormatted
		nearbyDestinationFormatted.ID = nearbyDestinationGotten.ID
		nearbyDestinationFormatted.Name = nearbyDestinationGotten.Name
		nearbyDestinationFormatted.Type = nearbyDestinationGotten.Type
		nearbyDestinationFormatted.GeoHash = nearbyDestinationGotten.GeoHash
		nearbyDestinationFormatted.ViewAmount = nearbyDestinationGotten.ViewAmount
		nearbyDestinationFormatted.Rating = nearbyDestinationGotten.Rating
		nearbyDestinationFormatted.Distance = distanceString

		// Tambahkan destinationFormatted ke allDestinationsFormatted
		allNearbyDestinationsFormatted = append(allNearbyDestinationsFormatted, nearbyDestinationFormatted)
	}

	return allNearbyDestinationsFormatted
}
