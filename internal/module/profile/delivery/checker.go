package delivery

import (
	"strconv"

	"github.com/banggibima/agile-backend/internal/module/profile/domain"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ProfileChecker struct{}

func NewProfileChecker() *ProfileChecker {
	return &ProfileChecker{}
}

func (c *ProfileChecker) Find(page, size int, sort, order string) error {
	check := validator.New()

	profile := &domain.FindRequest{
		Page:  strconv.Itoa(page),
		Size:  strconv.Itoa(size),
		Sort:  sort,
		Order: order,
	}

	if err := check.Struct(profile); err != nil {
		return err
	}

	return nil
}

func (c *ProfileChecker) FindByID(id uuid.UUID) error {
	check := validator.New()

	profile := &domain.FindByIDRequest{
		ID: id,
	}

	if err := check.Struct(profile); err != nil {
		return err
	}

	return nil
}

func (c *ProfileChecker) Save(payload *domain.Profile) error {
	check := validator.New()

	profile := &domain.SaveRequest{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Phone:     payload.Phone,
		UserID:    payload.UserID,
	}

	if err := check.Struct(profile); err != nil {
		return err
	}

	return nil
}

func (c *ProfileChecker) Edit(payload *domain.Profile) error {
	check := validator.New()

	profile := &domain.EditRequest{
		ID:        payload.ID,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Phone:     payload.Phone,
		UserID:    payload.UserID,
	}

	if err := check.Struct(profile); err != nil {
		return err
	}

	return nil
}

func (c *ProfileChecker) EditPartial(payload *domain.Profile) error {
	check := validator.New()

	profile := &domain.EditPartialRequest{
		ID:        payload.ID,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Phone:     payload.Phone,
		UserID:    payload.UserID,
	}

	if err := check.Struct(profile); err != nil {
		return err
	}

	return nil
}

func (c *ProfileChecker) Remove(payload *domain.Profile) error {
	check := validator.New()

	profile := &domain.RemoveRequest{
		ID: payload.ID,
	}

	if err := check.Struct(profile); err != nil {
		return err
	}

	return nil
}
