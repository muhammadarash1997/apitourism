package image

type Service interface {
	AddImage(imageInput ImageInput, fileName string) (Image, error)
	DeleteImageByUUID(id string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (this *service) AddImage(imageInput ImageInput, fileName string) (Image, error) {
	image := Image{}

	image.DestinationID = imageInput.DestinationID
	image.FileName = fileName

	imageAdded, err := this.repository.Save(image)
	if err != nil {
		return imageAdded, err
	}
	
	return imageAdded, nil
}

func (this *service) DeleteImageByUUID(id string) error {
	err := this.repository.DeleteByUUID(id)
	if err != nil {
		return err
	}

	return nil
}