package delivery

import (
	"strconv"

	"github.com/banggibima/agile-backend/internal/module/tag/domain"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type TagChecker struct{}

func NewTagChecker() *TagChecker {
	return &TagChecker{}
}

func (c *TagChecker) Find(page, size int, sort, order string) error {
	check := validator.New()

	tag := &domain.FindRequest{
		Page:  strconv.Itoa(page),
		Size:  strconv.Itoa(size),
		Sort:  sort,
		Order: order,
	}

	if err := check.Struct(tag); err != nil {
		return err
	}

	return nil
}

func (c *TagChecker) FindByID(id uuid.UUID) error {
	check := validator.New()

	tag := &domain.FindByIDRequest{
		ID: id,
	}

	if err := check.Struct(tag); err != nil {
		return err
	}

	return nil
}

func (c *TagChecker) Save(payload *domain.Tag) error {
	check := validator.New()

	tag := &domain.SaveRequest{
		Name: payload.Name,
	}

	if err := check.Struct(tag); err != nil {
		return err
	}

	return nil
}

func (c *TagChecker) Edit(payload *domain.Tag) error {
	check := validator.New()

	tag := &domain.EditRequest{
		ID:   payload.ID,
		Name: payload.Name,
	}

	if err := check.Struct(tag); err != nil {
		return err
	}

	return nil
}

func (c *TagChecker) EditPartial(payload *domain.Tag) error {
	check := validator.New()

	tag := &domain.EditPartialRequest{
		ID:   payload.ID,
		Name: payload.Name,
	}

	if err := check.Struct(tag); err != nil {
		return err
	}

	return nil
}

func (c *TagChecker) Remove(payload *domain.Tag) error {
	check := validator.New()

	tag := &domain.RemoveRequest{
		ID: payload.ID,
	}

	if err := check.Struct(tag); err != nil {
		return err
	}

	return nil
}
