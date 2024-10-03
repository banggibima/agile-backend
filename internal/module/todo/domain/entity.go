package domain

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
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
	Title   string `validate:"required"`
	Caption string `validate:"required"`
}

type EditRequest struct {
	ID      uuid.UUID `validate:"required"`
	Title   string    `validate:"required"`
	Caption string    `validate:"required"`
}

type EditPartialRequest struct {
	ID      uuid.UUID `validate:"required"`
	Title   string    `validate:"omitempty"`
	Caption string    `validate:"omitempty"`
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
	Data []*Todo `json:"data"`
}

type Detail struct {
	Data *Todo `json:"data"`
}

type Error struct {
	Error interface{} `json:"error"`
}
