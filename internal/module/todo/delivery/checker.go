package delivery

import (
	"strconv"

	"github.com/banggibima/agile-backend/internal/module/todo/domain"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type TodoChecker struct{}

func NewTodoChecker() *TodoChecker {
	return &TodoChecker{}
}

func (c *TodoChecker) Find(page, size int, sort, order string) error {
	check := validator.New()

	todo := &domain.FindRequest{
		Page:  strconv.Itoa(page),
		Size:  strconv.Itoa(size),
		Sort:  sort,
		Order: order,
	}

	if err := check.Struct(todo); err != nil {
		return err
	}

	return nil
}

func (c *TodoChecker) FindByID(id uuid.UUID) error {
	check := validator.New()

	todo := &domain.FindByIDRequest{
		ID: id,
	}

	if err := check.Struct(todo); err != nil {
		return err
	}

	return nil
}

func (c *TodoChecker) Save(payload *domain.Todo) error {
	check := validator.New()

	todo := &domain.SaveRequest{
		Title:   payload.Title,
		Caption: payload.Caption,
	}

	if err := check.Struct(todo); err != nil {
		return err
	}

	return nil
}

func (c *TodoChecker) Edit(payload *domain.Todo) error {
	check := validator.New()

	todo := &domain.EditRequest{
		ID:      payload.ID,
		Title:   payload.Title,
		Caption: payload.Caption,
	}

	if err := check.Struct(todo); err != nil {
		return err
	}

	return nil
}

func (c *TodoChecker) EditPartial(payload *domain.Todo) error {
	check := validator.New()

	todo := &domain.EditPartialRequest{
		ID:      payload.ID,
		Title:   payload.Title,
		Caption: payload.Caption,
	}

	if err := check.Struct(todo); err != nil {
		return err
	}

	return nil
}

func (c *TodoChecker) Remove(payload *domain.Todo) error {
	check := validator.New()

	todo := &domain.RemoveRequest{
		ID: payload.ID,
	}

	if err := check.Struct(todo); err != nil {
		return err
	}

	return nil
}
