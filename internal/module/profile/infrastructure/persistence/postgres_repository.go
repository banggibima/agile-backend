package persistence

import (
	"database/sql"
	"strconv"
	"strings"
	"time"

	"github.com/banggibima/agile-backend/internal/module/profile/domain"
	"github.com/google/uuid"
)

type ProfilePostgresRepository struct {
	DB *sql.DB
}

func NewProfilePostgresRepository(db *sql.DB) *ProfilePostgresRepository {
	return &ProfilePostgresRepository{
		DB: db,
	}
}

func (r *ProfilePostgresRepository) Count() (int, error) {
	total := 0
	query := "SELECT COUNT(*) FROM profiles"

	result := make(chan int)
	error := make(chan error)

	go func() {
		err := r.DB.QueryRow(query).Scan(&total)
		if err != nil {
			error <- err
		}
		result <- total
	}()

	select {
	case res := <-result:
		return res, nil
	case err := <-error:
		return 0, err
	}
}

func (r *ProfilePostgresRepository) Find(offset, limit int, sort, order string) ([]*domain.Profile, error) {
	query := "SELECT id, first_name, last_name, email, phone, user_id, created_at, updated_at FROM profiles"

	if sort != "" && order != "" {
		query += " ORDER BY " + sort + " " + strings.ToUpper(order)
	} else {
		query += " ORDER BY created_at DESC"
	}

	if limit > 0 && offset >= 0 {
		query += " LIMIT $1 OFFSET $2"

		result := make(chan []*domain.Profile)
		error := make(chan error)

		go func() {
			rows, err := r.DB.Query(query, limit, offset)
			if err != nil {
				error <- err
			}
			defer rows.Close()

			profiles := []*domain.Profile{}
			for rows.Next() {
				profile := domain.Profile{}
				if err := rows.Scan(
					&profile.ID,
					&profile.FirstName,
					&profile.LastName,
					&profile.Email,
					&profile.Phone,
					&profile.UserID,
					&profile.CreatedAt,
					&profile.UpdatedAt,
				); err != nil {
					error <- err
				}
				profiles = append(profiles, &profile)
			}

			if err := rows.Err(); err != nil {
				error <- err
			}

			result <- profiles
		}()

		select {
		case res := <-result:
			return res, nil
		case err := <-error:
			return nil, err
		}
	} else {

		result := make(chan []*domain.Profile)
		error := make(chan error)

		go func() {
			rows, err := r.DB.Query(query)
			if err != nil {
				error <- err
			}
			defer rows.Close()

			profiles := []*domain.Profile{}
			for rows.Next() {
				profile := domain.Profile{}
				if err := rows.Scan(
					&profile.ID,
					&profile.FirstName,
					&profile.LastName,
					&profile.Email,
					&profile.Phone,
					&profile.UserID,
					&profile.CreatedAt,
					&profile.UpdatedAt,
				); err != nil {
					error <- err
				}
				profiles = append(profiles, &profile)
			}

			if err := rows.Err(); err != nil {
				error <- err
			}

			result <- profiles
		}()

		select {
		case res := <-result:
			return res, nil
		case err := <-error:
			return nil, err
		}
	}
}

func (r *ProfilePostgresRepository) FindByID(id uuid.UUID) (*domain.Profile, error) {
	query := "SELECT id, first_name, last_name, email, phone, user_id, created_at, updated_at FROM profiles WHERE id = $1"

	result := make(chan *domain.Profile)
	error := make(chan error)

	go func() {
		profile := domain.Profile{}
		err := r.DB.QueryRow(query, id).Scan(
			&profile.ID,
			&profile.FirstName,
			&profile.LastName,
			&profile.Email,
			&profile.Phone,
			&profile.UserID,
			&profile.CreatedAt,
			&profile.UpdatedAt,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				result <- nil
			}
			error <- err
		}
		result <- &profile
	}()

	select {
	case res := <-result:
		return res, nil
	case err := <-error:
		return nil, err
	}
}

func (r *ProfilePostgresRepository) Save(payload *domain.Profile) error {
	query := "INSERT INTO profiles (id, first_name, last_name, email, phone, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, first_name, last_name, email, phone, user_id, created_at, updated_at"

	payload.ID = uuid.New()
	payload.CreatedAt = time.Now()
	payload.UpdatedAt = time.Now()

	result := make(chan error)
	error := make(chan error)

	go func() {
		err := r.DB.QueryRow(query,
			payload.ID,
			payload.FirstName,
			payload.LastName,
			payload.Email,
			payload.Phone,
			payload.UserID,
			payload.CreatedAt,
			payload.UpdatedAt,
		).Scan(
			&payload.ID,
			&payload.FirstName,
			&payload.LastName,
			&payload.Email,
			&payload.Phone,
			&payload.UserID,
			&payload.CreatedAt,
			&payload.UpdatedAt,
		)
		if err != nil {
			error <- err
		}
		result <- nil
	}()

	select {
	case res := <-result:
		return res
	case err := <-error:
		return err
	}
}

func (r *ProfilePostgresRepository) Edit(payload *domain.Profile) error {
	query := "UPDATE profiles SET first_name = $1, last_name = $2, email = $3, phone = $4, user_id = $5, updated_at = $6 WHERE id = $7 RETURNING id, first_name, last_name, email, phone, user_id, created_at, updated_at"

	payload.UpdatedAt = time.Now()

	result := make(chan error)
	error := make(chan error)

	go func() {
		err := r.DB.QueryRow(query,
			payload.FirstName,
			payload.LastName,
			payload.Email,
			payload.Phone,
			payload.UserID,
			payload.UpdatedAt,
			payload.ID,
		).Scan(
			&payload.ID,
			&payload.FirstName,
			&payload.LastName,
			&payload.Email,
			&payload.Phone,
			&payload.UserID,
			&payload.CreatedAt,
			&payload.UpdatedAt,
		)
		if err != nil {
			error <- err
		}
		result <- nil
	}()

	select {
	case res := <-result:
		return res
	case err := <-error:
		return err
	}
}

func (r *ProfilePostgresRepository) EditPartial(payload *domain.Profile) error {
	query := "UPDATE profiles SET "
	args := []interface{}{}
	clauses := []string{}

	if payload.FirstName != "" {
		clauses = append(clauses, "first_name = $"+strconv.Itoa(len(args)+1))
		args = append(args, payload.FirstName)
	}

	if payload.LastName != "" {
		clauses = append(clauses, "last_name = $"+strconv.Itoa(len(args)+1))
		args = append(args, payload.LastName)
	}

	if payload.Email != "" {
		clauses = append(clauses, "email = $"+strconv.Itoa(len(args)+1))
		args = append(args, payload.Email)
	}

	if payload.Phone != "" {
		clauses = append(clauses, "phone = $"+strconv.Itoa(len(args)+1))
		args = append(args, payload.Phone)
	}

	if payload.UserID != uuid.Nil {
		clauses = append(clauses, "user_id = $"+strconv.Itoa(len(args)+1))
		args = append(args, payload.UserID)
	}

	payload.UpdatedAt = time.Now()

	clauses = append(clauses, "updated_at = $"+strconv.Itoa(len(args)+1))
	args = append(args, payload.UpdatedAt)

	if len(clauses) == 0 {
		return nil
	}

	query += strings.Join(clauses, ", ") + " WHERE id = $" + strconv.Itoa(len(args)+1) + " RETURNING id, first_name, last_name, email, phone, user_id, created_at, updated_at"
	args = append(args, payload.ID)

	result := make(chan error)
	error := make(chan error)

	go func() {
		err := r.DB.QueryRow(query, args...).Scan(
			&payload.ID,
			&payload.FirstName,
			&payload.LastName,
			&payload.Email,
			&payload.Phone,
			&payload.UserID,
			&payload.CreatedAt,
			&payload.UpdatedAt,
		)
		if err != nil {
			error <- err
		}
		result <- nil
	}()

	select {
	case res := <-result:
		return res
	case err := <-error:
		return err
	}
}

func (r *ProfilePostgresRepository) Remove(payload *domain.Profile) error {
	query := "DELETE FROM profiles WHERE id = $1"

	result := make(chan error)
	error := make(chan error)

	go func() {
		_, err := r.DB.Exec(query, payload.ID)
		if err != nil {
			error <- err
		}
		result <- nil
	}()

	select {
	case res := <-result:
		return res
	case err := <-error:
		return err
	}
}
