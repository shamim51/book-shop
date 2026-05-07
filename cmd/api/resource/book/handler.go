package book

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	repository *Repository
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		repository: NewRepository(db),
	}
}

func (service *Service) List(w http.ResponseWriter, r *http.Request) {
	books, err := service.repository.List()
	if err != nil {
		//TODO: need to use appropriate and structured error message
		return
	}

	if len(books) == 0 {
		fmt.Fprintf(w, "[]")
		return
	}

	if err := json.NewEncoder(w).Encode(books); err != nil {
		//TODO: need to use appropriate and structured error message
		return
	}
}

func (service *Service) Create(w http.ResponseWriter, r *http.Request) {
	bookRequest := &BookRequest{}
	if err := json.NewDecoder(r.Body).Decode(bookRequest); err != nil {
		//TODO: need to use appropriate and structured error message
		return
	}

	newBook := bookRequest.ToBook()
	newBook.ID = uuid.New()

	_, err := service.repository.Create(newBook)

	if err != nil {
		//TODO: need to use appropriate and structured error message
		return
	}

	w.WriteHeader(http.StatusCreated)

}
func (service *Service) Read(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		//TODO: need to use appropriate and structured error message
		return
	}

	book, err := service.repository.Read(id)

	if err != nil {
		//TODO: need to use appropriate and structured error message
		return
	}

	dto := book.ToDto()

	if err := json.NewEncoder(w).Encode(dto); err != nil {
		//TODO: need to use appropriate and structured error message
		return
	}

}

func (service *Service) Update(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		//TODO: need to use appropriate and structured error message
		return
	}

	bookRequest := &BookRequest{}
	if err := json.NewDecoder(r.Body).Decode(bookRequest); err != nil {
		return
	}

	book := bookRequest.ToBook()
	book.ID = id

	rows, err := service.repository.Update(book)
	if err != nil {
		//TODO: need to use appropriate and structured error message
		return
	}

	if rows == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

}
func (service *Service) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		//TODO: need to use appropriate and structured error message
		return
	}

	rows, err := service.repository.Delete(id)
	if err != nil {
		//TODO: need to use appropriate and structured error message
		return
	}

	if rows == 0 {
		//TODO: need to use appropriate and structured error message
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
