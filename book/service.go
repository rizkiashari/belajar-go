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

func (s *service) Create(bookInput BookInput) (Book, error) {

	price, _ := bookInput.Price.Int64()
	rating, _ := bookInput.Rating.Int64()
	book := Book{
		Title:       bookInput.Title,
		Price:       int(price),
		Description: bookInput.Description,
		Rating:      int(rating),
	}

	newBook, err := s.repository.Create(book)

	return newBook, err
}

func (s *service) Update(ID int, bookInput BookInput) (Book, error) {

	book, _ := s.repository.FindByID(ID)

	price, _ := bookInput.Price.Int64()
	rating, _ := bookInput.Rating.Int64()

	book.Title = bookInput.Title
	book.Price = int(price)
	book.Description = bookInput.Description
	book.Rating = int(rating)

	newBook, err := s.repository.Update(book)

	return newBook, err
}

func (s *service) DeleteBook(ID int) (Book, error) {

	book, _ := s.repository.FindByID(ID)

	newBook, err := s.repository.Delete(book)

	return newBook, err
}
