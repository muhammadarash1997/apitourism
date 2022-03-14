// Package classification Tourism API.
//
// This is a API for tourism to search destinations
//
//   Schemes: http
//   Host: https://apitourism.herokuapp.com
//   BasePath: /swagger/
//   Version: 1.0.0
//   Contact: muhammadarash1997@gmail.com
//
//   Consumes:
//   - application/json
//
//   Produces:
//   - application/json
//
//   SecurityDefinitions:
//   Bearer:
//    type: apiKey
//    name: Authorization
//    in: header
//
// swagger:meta
package docs

import (
	"apitourism/bookmark"
	"apitourism/destination"
	"apitourism/helper"
	"apitourism/image"
	"apitourism/rating"
	"apitourism/user"
	"apitourism/view"
	"bytes"
)

// Success adding new destination
// swagger:response destinationAdded200
type destinationAdded200 struct {
	// in: Body
	Body struct {
		Meta helper.Meta `json:"meta"`
		Data destination.DestinationAddedFormatter `json:"data"`
	}
}

// Success deleting a destination
// swagger:response destinationDeleted200
type destinationDeleted200 struct {
	// in: Body
	Body struct {
		Meta helper.Meta `json:"meta"`
		Data struct {
		} `json:"data"`
	}
}

// Success adding an image
// swagger:response imageAdded200
type imageAdded200 struct {
	// in: Body
	Body struct {
		Meta helper.Meta `json:"meta"`
		Data image.ImageFormatter `json:"data"`
	}
}

// Success deleting an image
// swagger:response imageDeleted200
type imageDeleted200 struct {
	// in: Body
	Body struct {
		Meta helper.Meta `json:"meta"`
		Data struct {
		} `json:"data"`
	}
}

// Success registering an account
// swagger:response userRegistered200
type userRegistered200 struct {
	// in: Body
	Body struct {
		Meta helper.Meta `json:"meta"`
		Data user.UserRegisteredFormatter `json:"data"`
	}
}

// Success logging in an account
// swagger:response userLogged200
type userLogged200 struct {
	// in: Body
	Body struct {
		Meta helper.Meta `json:"meta"`
		Data user.UserLoggedFormatter `json:"data"`
	}
}

// Success deleting a user
// swagger:response userDeleted200
type userDeleted200 struct {
	// in: Body
	Body struct {
		Meta helper.Meta `json:"meta"`
		Data struct {
		} `json:"data"`
	}
}

// Success adding a bookmark
// swagger:response bookmarkAdded200
type bookmarkAdded200 struct {
	// in: Body
	Body struct {
		Meta helper.Meta `json:"meta"`
		Data bookmark.BookmarkFormatter `json:"data"`
	}
}

// Success adding a view
// swagger:response viewAdded200
type viewAdded200 struct {
	// in: Body
	Body struct {
		Meta helper.Meta `json:"meta"`
		Data view.ViewFormatter `json:"data"`
	}
}

// Success adding rating
// swagger:response ratingAdded200
type ratingAdded200 struct {
	// in: Body
	Body struct {
		Meta helper.Meta `json:"meta"`
		Data rating.RatingFormatter `json:"data"`
	}
}

// Success deleting bookmark
// swagger:response bookmarkDeleted200
type bookmarkDeleted200 struct {
	// in: Body
	Body struct {
		Meta helper.Meta `json:"meta"`
		Data struct {
		} `json:"data"`
	}
}

// Success getting destinations
// swagger:response destinationsGotten200
type destinationsGotten200 struct {
	// in: Body
	Body struct {
		Meta helper.Meta `json:"meta"`
		Data []destination.DestinationGottenFormatter `json:"data"`
	}
}

// Success getting nearby destinations
// swagger:response nearbyDestinationsGotten200
type nearbyDestinationsGotten200 struct {
	// in: Body
	Body struct {
		Meta helper.Meta `json:"meta"`
		Data []destination.NearbyDestinationGottenFormatter `json:"data"`
	}
}

// swagger:response errorResponse
type errorResponse struct {
	// in: Body
	Body struct {
		Meta helper.Meta `json:"meta"`
		Data struct {
			Errors string `json:"errors"`
		} `json:"data"`
	}
}

// swagger:parameters addDestination
type addDestinationParams struct {
	// Destination object that needs to be added to the store
	// in: body
	// required: true
	Body destination.DestinationInput
}

// swagger:parameters addImageByUUID
type addImageByUUIDParams struct {
	// Destination uuid to add image of destination
	// in: path
	// required: true
	ID string `json:"destinationUUID"`

	// Image file of destination
	// in: formData
	// swagger:file
	ImageFile *bytes.Buffer `json:"imageFile"`
}

// swagger:parameters deleteUserByUUID
type deleteUserByUUIDParams struct {
	// The user uuid that needs to be deleted
	// in: path
	// required: true
	ID string `json:"userUUID"`
}

// swagger:parameters deleteDestinationByUUID
type deleteDestinationByUUIDParams struct {
	// The destination uuid that needs to be deleted
	// in: path
	// required: true
	ID string `json:"destinationUUID"`
}

// swagger:parameters deleteImageByUUID
type deleteImageByUUIDParams struct {
	// The image uuid that needs to be deleted
	// in: path
	// required: true
	ID string `json:"imageUUID"`
}

// swagger:parameters registerUser
type registerUserParams struct {
	// User object that needs to be registered
	// in: body
	// required: true
	Body struct {
		Name string `json:"name"`
		Email string `json:"email"`
		Password string `json:"password"`
	}
}

// swagger:parameters loginUser
type loginUserParams struct {
	// User object that needs to be logged in
	// in: body
	// required: true
	Body user.LoginInput
}

// swagger:parameters addBookmark
type addBookmarkParams struct {
	// Bookmark object that needs to be added
	// in: body
	// required: true
	Body bookmark.BookmarkInput
}

// swagger:parameters addView
type addViewParams struct {
	// View object that needs to be added
	// in: body
	// required: true
	Body view.ViewInput
}

// swagger:parameters addRating
type addRatingParams struct {
	// Rating object that needs to be added
	// in: body
	// required: true
	Body struct {
		DestinationID string `json:"destination_id"`
		UserID string `json:"user_id"`
		// enum: 1, 2, 3, 4, 5
		Rate int `json:"rate"`
	}
}

// swagger:parameters deleteBookmarkByUUID
type deleteBookmarkByUUIDParams struct {
	// The bookmark uuid that needs to be deleted
	// in: path
	// required: true
	ID string `json:"bookmarkUUID"`
}

// swagger:parameters getAllDestinations
type getAllDestinationsParams struct {
	// User coordinate for knowing user location, ex = -8.5797192,116.0987078
	// in: path
	// required: true
	UserCoordinate string `json:"userCoordinate"`

	// The numbers of items to return
	// in: query
	// required: true
	Limit int `json:"limit"`

	// The page of items to return
	// in: query
	// required: true
	Page int `json:"page"`
}

// swagger:parameters getNearbyDestinations
type getNearbyDestinationsParams struct {
	// User coordinate for knowing user location, ex = -8.5797192,116.0987078
	// in: path
	// required: true
	UserCoordinate string `json:"userCoordinate"`

	// The type of destinations as a parameter for finding
	// in: query
	// required: true
	// enum: all, beach, natural, towns, heritage
	Type string `json:"type"`

	// Use rating as a parameter for finding
	// in: query
	// required: true
	Rating bool `json:"rating"`

	// Use popularity as a parameter for finding
	// in: query
	// required: true
	Popularity bool `json:"popularity"`
}