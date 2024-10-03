package delivery

import (
	"strconv"

	"github.com/banggibima/agile-backend/internal/module/post/domain"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type PostChecker struct{}

func NewPostChecker() *PostChecker {
	return &PostChecker{}
}

func (c *PostChecker) Find(page, size int, sort, order string) error {
	check := validator.New()

	post := &domain.FindRequest{
		Page:  strconv.Itoa(page),
		Size:  strconv.Itoa(size),
		Sort:  sort,
		Order: order,
	}

	if err := check.Struct(post); err != nil {
		return err
	}

	return nil
}

func (c *PostChecker) FindByID(id uuid.UUID) error {
	check := validator.New()

	post := &domain.FindByIDRequest{
		ID: id,
	}

	if err := check.Struct(post); err != nil {
		return err
	}

	return nil
}

func (c *PostChecker) Save(payload *domain.Post) error {
	check := validator.New()

	post := &domain.SaveRequest{
		Title:   payload.Title,
		Content: payload.Content,
		UserID:  payload.UserID,
	}

	if err := check.Struct(post); err != nil {
		return err
	}

	return nil
}

func (c *PostChecker) Edit(payload *domain.Post) error {
	check := validator.New()

	post := &domain.EditRequest{
		ID:      payload.ID,
		Title:   payload.Title,
		Content: payload.Content,
		UserID:  payload.UserID,
	}

	if err := check.Struct(post); err != nil {
		return err
	}

	return nil
}

func (c *PostChecker) EditPartial(payload *domain.Post) error {
	check := validator.New()

	post := &domain.EditPartialRequest{
		ID:      payload.ID,
		Title:   payload.Title,
		Content: payload.Content,
		UserID:  payload.UserID,
	}

	if err := check.Struct(post); err != nil {
		return err
	}

	return nil
}

func (c *PostChecker) Remove(payload *domain.Post) error {
	check := validator.New()

	post := &domain.RemoveRequest{
		ID: payload.ID,
	}

	if err := check.Struct(post); err != nil {
		return err
	}

	return nil
}
