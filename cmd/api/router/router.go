package router

import (
	"book-shop/cmd/api/resource/book"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB, validator *validator.Validate) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/v1", func(r chi.Router) {
		bookService := book.NewService(db, validator)

		r.Get("/books", bookService.List)
		r.Post("/books", bookService.Create)
		r.Get("/books/{id}", bookService.Read)
		r.Put("/books/{id}", bookService.Update)
		r.Delete("/books/{id}", bookService.Delete)

	})

	return r
}
