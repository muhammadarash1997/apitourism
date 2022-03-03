package rating

type Service interface {
	AddRating(ratingInput RatingInput) (Rating, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (this *service) AddRating(ratingInput RatingInput) (Rating, error) {
	rating := Rating{}

	rating.DestinationID = ratingInput.DestinationID
	rating.UserID = ratingInput.UserID
	rating.Rate = ratingInput.Rate

	ratingAdded, err := this.repository.Add(rating)
	if err != nil {
		return ratingAdded, err
	}

	return ratingAdded, nil
}
