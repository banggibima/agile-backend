package domain

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
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
	FirstName string    `validate:"required"`
	LastName  string    `validate:"required"`
	Email     string    `validate:"required,email"`
	Phone     string    `validate:"required"`
	UserID    uuid.UUID `validate:"required"`
}

type EditRequest struct {
	ID        uuid.UUID `validate:"required"`
	FirstName string    `validate:"required"`
	LastName  string    `validate:"required"`
	Email     string    `validate:"required,email"`
	Phone     string    `validate:"required"`
	UserID    uuid.UUID `validate:"required"`
}

type EditPartialRequest struct {
	ID        uuid.UUID `validate:"required"`
	FirstName string    `validate:"omitempty"`
	LastName  string    `validate:"omitempty"`
	Email     string    `validate:"omitempty,email"`
	Phone     string    `validate:"omitempty"`
	UserID    uuid.UUID `validate:"omitempty"`
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
	Meta *Meta      `json:"meta"`
	Data []*Profile `json:"data"`
}

type Detail struct {
	Data *Profile `json:"data"`
}

type Error struct {
	Error interface{} `json:"error"`
}
