package delivery

import (
	"strconv"

	"github.com/banggibima/agile-backend/internal/module/user/domain"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserChecker struct{}

func NewUserChecker() *UserChecker {
	return &UserChecker{}
}

func (c *UserChecker) Find(page, size int, sort, order string) error {
	check := validator.New()

	user := &domain.FindRequest{
		Page:  strconv.Itoa(page),
		Size:  strconv.Itoa(size),
		Sort:  sort,
		Order: order,
	}

	if err := check.Struct(user); err != nil {
		return err
	}

	return nil
}

func (c *UserChecker) FindByID(id uuid.UUID) error {
	check := validator.New()

	user := &domain.FindByIDRequest{
		ID: id,
	}

	if err := check.Struct(user); err != nil {
		return err
	}

	return nil
}

func (c *UserChecker) Save(payload *domain.User) error {
	check := validator.New()

	user := &domain.SaveRequest{
		Username: payload.Username,
		Password: payload.Password,
		Role:     payload.Role,
		Status:   payload.Status,
	}

	if err := check.Struct(user); err != nil {
		return err
	}

	return nil
}

func (c *UserChecker) Edit(payload *domain.User) error {
	check := validator.New()

	user := &domain.EditRequest{
		ID:       payload.ID,
		Username: payload.Username,
		Password: payload.Password,
		Role:     payload.Role,
		Status:   payload.Status,
	}

	if err := check.Struct(user); err != nil {
		return err
	}

	return nil
}

func (c *UserChecker) EditPartial(payload *domain.User) error {
	check := validator.New()

	user := &domain.EditPartialRequest{
		ID:       payload.ID,
		Username: payload.Username,
		Password: payload.Password,
		Role:     payload.Role,
		Status:   payload.Status,
	}

	if err := check.Struct(user); err != nil {
		return err
	}

	return nil
}

func (c *UserChecker) Remove(payload *domain.User) error {
	check := validator.New()

	user := &domain.RemoveRequest{
		ID: payload.ID,
	}

	if err := check.Struct(user); err != nil {
		return err
	}

	return nil
}
