package bookmark

type Service interface {
	AddBookmark(bookmarkInput BookmarkInput) (Bookmark, error)
	DeleteBookmarkByUUID(id string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (this *service) AddBookmark(bookmarkInput BookmarkInput) (Bookmark, error) {
	bookmark := Bookmark{}

	bookmark.UserID = bookmarkInput.UserID
	bookmark.DestinationID = bookmarkInput.DestinationID

	bookmarkAdded, err := this.repository.Add(bookmark)
	if err != nil {
		return bookmarkAdded, err
	}

	return bookmarkAdded, nil
}

func (this *service) DeleteBookmarkByUUID(id string) error {
	err := this.repository.DeleteByUUID(id)
	if err != nil {
		return err
	}

	return nil
}