package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookInput BookInput) (Book, error)
	Update(ID int, bookInput BookInput) (Book, error)
	Delete(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
}

func (s *service) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	return book, err
}

func (s *service) Create(bookRequest BookInput) (Book, error) {
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()

	book := Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Description: bookRequest.Description,
		Rating:      int(rating),
	}

	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Update(ID int, bookInput BookInput) (Book, error) {
	book, err := s.FindByID(ID)

	price, _ := bookInput.Price.Int64()
	rating, _ := bookInput.Rating.Int64()

	book.Title = bookInput.Title
	book.Description = bookInput.Description
	book.Price = int(price)
	book.Rating = int(rating)

	s.repository.Update(book)
	return book, err
}

func (s *service) Delete(ID int) (Book, error) {
	book, _ := s.repository.FindByID(ID)

	newBook, err := s.repository.Delete(book)

	return newBook, err
}
