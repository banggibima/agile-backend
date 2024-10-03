package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	Status    string    `json:"status"`
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
	Username string `validate:"required"`
	Password string `validate:"required"`
	Role     string `validate:"required"`
	Status   string `validate:"required"`
}

type EditRequest struct {
	ID       uuid.UUID `validate:"required"`
	Username string    `validate:"required"`
	Password string    `validate:"required"`
	Role     string    `validate:"required"`
	Status   string    `validate:"required"`
}

type EditPartialRequest struct {
	ID       uuid.UUID `validate:"required"`
	Username string    `validate:"omitempty"`
	Password string    `validate:"omitempty"`
	Role     string    `validate:"omitempty"`
	Status   string    `validate:"omitempty"`
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
	Data []*User `json:"data"`
}

type Detail struct {
	Data *User `json:"data"`
}

type Error struct {
	Error interface{} `json:"error"`
}
