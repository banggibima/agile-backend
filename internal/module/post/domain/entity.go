package domain

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FindRequest struct {
	Page  string `validate:"omitempty"`
	Size  string `validate:"omitempty"`
	Sort  string `validate:"omitempty"`
	Order string `validate:"omitempty"`
}

type FindByIDRequest struct {
	ID uuid.UUID `validate:"required"`
}

type SaveRequest struct {
	Title   string    `validate:"required"`
	Content string    `validate:"required"`
	UserID  uuid.UUID `validate:"required"`
}

type EditRequest struct {
	ID      uuid.UUID `validate:"required"`
	Title   string    `validate:"required"`
	Content string    `validate:"required"`
	UserID  uuid.UUID `validate:"required"`
}

type EditPartialRequest struct {
	ID      uuid.UUID `validate:"required"`
	Title   string    `validate:"omitempty"`
	Content string    `validate:"omitempty"`
	UserID  uuid.UUID `validate:"omitempty"`
}

type RemoveRequest struct {
	ID uuid.UUID `validate:"required"`
}

type Meta struct {
	Page  int    `json:"page"`
	Size  int    `json:"size"`
	Count int    `json:"count"`
	Total int    `json:"total"`
	Sort  string `json:"sort"`
	Order string `json:"order"`
}

type List struct {
	Meta *Meta   `json:"meta"`
	Data []*Post `json:"data"`
}

type Detail struct {
	Data *Post `json:"data"`
}

type Error struct {
	Error interface{} `json:"error"`
}
