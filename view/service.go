package view

type Service interface {
	AddView(viewInput ViewInput) (View, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (this *service) AddView(viewInput ViewInput) (View, error) {
	view := View{}

	view.DestinationID = viewInput.DestinationID

	viewAdded, err := this.repository.Add(view)
	if err != nil {
		return viewAdded, err
	}

	return viewAdded, nil
}