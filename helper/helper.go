package helper

import (
	. "math"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/pierrre/geohash"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	metaOutput := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	responseOutput := Response{
		Meta: metaOutput,
		Data: data,
	}

	return responseOutput
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func CoordinateParser(coordinate string) (float64, float64) {
	coordinateSplited := strings.Split(coordinate, ",")
	latitude, _ := strconv.ParseFloat(coordinateSplited[0], 64)
	longitude, _ := strconv.ParseFloat(coordinateSplited[1], 64)

	return latitude, longitude
}

func DecodeGeohash(geohashEncoded string) (float64, float64) {
	geohashBox, _ := geohash.Decode(geohashEncoded)
	geohashDecoded := geohashBox.Round()
	latitudeDecoded := geohashDecoded.Lat
	longitudeDecoded := geohashDecoded.Lon

	return latitudeDecoded, longitudeDecoded
}

func CalculateDistance(latitude1 float64, longitude1 float64, latitude2 float64, longitude2 float64) float64 {
	// Distance 2 Points
	distanceInMeter := 6371000 * Acos(Sin(latitude1*Pi/180)*Sin(latitude2*Pi/180)+Cos(latitude1*Pi/180)*Cos(latitude2*Pi/180)*Cos(longitude2*Pi/180-longitude1*Pi/180)) // Hasil dalam meter
	distanceInKiloMeter := float64(int(distanceInMeter*1)/100) / 10                                                                                                     // Hasil dalam kilometer

	return distanceInKiloMeter
}

func EncodeGeohash(coordinate string) string {
	latitude, longitude := CoordinateParser(coordinate)
	geohashEncoded := geohash.EncodeAuto(latitude, longitude)

	return geohashEncoded
}
