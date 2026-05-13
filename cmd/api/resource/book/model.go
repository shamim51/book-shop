package book

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID            uuid.UUID `gorm:"primaryKey"`
	Title         string
	Author        string
	PublishedDate time.Time
	ImageUrl      string
	Description   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

type Books []*Book

type DTO struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedDate string `json:"published_date"`
	ImageURL      string `json:"image_url"`
	Description   string `json:"description"`
}

type BookRequest struct {
	Title         string `json:"title" validate:"required,max=20"`
	Author        string `json:"author" validate:"required,alphaspace,max=20"`
	PublishedDate string `json:"published_date" validate:"required,datetime=2026-03-04"`
	ImageURL      string `json:"image_url" validate:"url"`
	Description   string `json:"description" validate:"required"`
}

func (f *BookRequest) ToBook() *Book {
	pubDate, _ := time.Parse("2006-01-02", f.PublishedDate)

	return &Book{
		Title:         f.Title,
		Author:        f.Author,
		PublishedDate: pubDate,
		ImageUrl:      f.ImageURL,
		Description:   f.Description,
	}
}

func (b *Book) ToDto() *DTO {
	return &DTO{
		ID:            b.ID.String(),
		Title:         b.Title,
		Author:        b.Author,
		PublishedDate: b.PublishedDate.Format("2006-01-02"),
		ImageURL:      b.ImageUrl,
		Description:   b.Description,
	}
}

func (bs Books) ToDto() []*DTO {
	dtos := make([]*DTO, len(bs))
	for i, v := range bs {
		dtos[i] = v.ToDto()
	}

	return dtos
}
