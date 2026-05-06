package book

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type Service struct {
	repository *Repository
}

func New(db *gorm.DB) *Service {
	return &Service{
		repository: NewRepository(db),
	}
}

func (service *Service) List(w http.ResponseWriter, r *http.Request) {
	boooks, err := service.repository.List()
	if err != nil {
		//TODO: need to use appropriate and structured error message
		return
	}

	if len(boooks) == 0 {
		fmt.Fprintf(w, "[]")
		return
	}

	if err := json.NewEncoder(w).Encode(boooks); err != nil {
		//TODO: need to use appropriate and structured error message
		return
	}
}

func (service *Service) Create(w http.ResponseWriter, r *http.Request) {

}
func (service *Service) Read(w http.ResponseWriter, r *http.Request) {

}

func (service *Service) Update(w http.ResponseWriter, r *http.Request) {

}
func (service *Service) Delete(w http.ResponseWriter, r *http.Request) {

}
