package book

import (
	"encoding/json"
	"fmt"
	"net/http"

	e "book-shop/cmd/api/resource/common/err"

	validatorUtil "book-shop/util/validator"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	repository *Repository
	validator  *validator.Validate
}

func NewService(db *gorm.DB, v *validator.Validate) *Service {
	return &Service{
		repository: NewRepository(db),
		validator:  v,
	}
}

func (service *Service) List(w http.ResponseWriter, r *http.Request) {
	books, err := service.repository.List()
	if err != nil {
		e.ServerError(w, e.RespDBDataAccessFailure)
		return
	}

	if len(books) == 0 {
		fmt.Fprintf(w, "[]")
		return
	}

	if err := json.NewEncoder(w).Encode(books); err != nil {
		e.ServerError(w, e.RespJSONEncodeFailure)
		return
	}
}

func (service *Service) Create(w http.ResponseWriter, r *http.Request) {
	bookRequest := &BookRequest{}
	if err := json.NewDecoder(r.Body).Decode(bookRequest); err != nil {
		e.ServerError(w, e.RespJSONDecodeFailure)
		return
	}

	if err := service.validator.Struct(bookRequest); err != nil {
		respBody, err := json.Marshal(validatorUtil.ToErrResponse(err))
		if err != nil {
			e.ServerError(w, e.RespJSONEncodeFailure)
			return
		}

		e.ValidationErrors(w, respBody)
		return
	}
	newBook := bookRequest.ToBook()
	newBook.ID = uuid.New()

	_, err := service.repository.Create(newBook)

	if err != nil {
		e.ServerError(w, e.RespDBDataInsertFailure)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
func (service *Service) Read(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}

	book, err := service.repository.Read(id)

	if err != nil {
		e.ServerError(w, e.RespDBDataAccessFailure)
		return
	}

	dto := book.ToDto()

	if err := json.NewEncoder(w).Encode(dto); err != nil {
		e.ServerError(w, e.RespJSONEncodeFailure)
		return
	}

}

func (service *Service) Update(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}

	bookRequest := &BookRequest{}
	if err := json.NewDecoder(r.Body).Decode(bookRequest); err != nil {
		e.ServerError(w, e.RespJSONDecodeFailure)
		return
	}

	if err := service.validator.Struct(bookRequest); err != nil {
		respBody, err := json.Marshal(validatorUtil.ToErrResponse(err))
		if err != nil {
			e.ServerError(w, e.RespJSONEncodeFailure)
			return
		}

		e.ValidationErrors(w, respBody)
		return
	}

	book := bookRequest.ToBook()
	book.ID = id

	rows, err := service.repository.Update(book)
	if err != nil {
		e.ServerError(w, e.RespDBDataUpdateFailure)
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
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}

	rows, err := service.repository.Delete(id)
	if err != nil {
		//TODO: need to use appropriate and structured error message
		return
	}

	if rows == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
