package persistence

import (
	"database/sql"
	"strconv"
	"strings"
	"time"

	"github.com/banggibima/agile-backend/internal/module/user/domain"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserPostgresRepository struct {
	DB *sql.DB
}

func NewUserPostgresRepository(db *sql.DB) *UserPostgresRepository {
	return &UserPostgresRepository{
		DB: db,
	}
}

func (r *UserPostgresRepository) Count() (int, error) {
	total := 0
	query := "SELECT COUNT(*) FROM users"

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

func (r *UserPostgresRepository) Find(offset, limit int, sort, order string) ([]*domain.User, error) {
	query := "SELECT id, username, password, role, status, created_at, updated_at FROM users"

	if sort != "" && order != "" {
		query += " ORDER BY " + sort + " " + strings.ToUpper(order)
	} else {
		query += " ORDER BY created_at DESC"
	}

	if limit > 0 && offset >= 0 {
		query += " LIMIT $1 OFFSET $2"

		result := make(chan []*domain.User)
		error := make(chan error)

		go func() {
			rows, err := r.DB.Query(query, limit, offset)
			if err != nil {
				error <- err
			}
			defer rows.Close()

			users := []*domain.User{}
			for rows.Next() {
				user := domain.User{}
				if err := rows.Scan(
					&user.ID,
					&user.Username,
					&user.Password,
					&user.Role,
					&user.Status,
					&user.CreatedAt,
					&user.UpdatedAt,
				); err != nil {
					error <- err
				}
				users = append(users, &user)
			}

			if err := rows.Err(); err != nil {
				error <- err
			}

			result <- users
		}()

		select {
		case res := <-result:
			return res, nil
		case err := <-error:
			return nil, err
		}
	} else {
		result := make(chan []*domain.User)
		error := make(chan error)

		go func() {
			rows, err := r.DB.Query(query)
			if err != nil {
				error <- err
			}
			defer rows.Close()

			users := []*domain.User{}
			for rows.Next() {
				user := domain.User{}
				if err := rows.Scan(
					&user.ID,
					&user.Username,
					&user.Password,
					&user.Role,
					&user.Status,
					&user.CreatedAt,
					&user.UpdatedAt,
				); err != nil {
					error <- err
				}
				users = append(users, &user)
			}

			if err := rows.Err(); err != nil {
				error <- err
			}

			result <- users
		}()

		select {
		case res := <-result:
			return res, nil
		case err := <-error:
			return nil, err
		}
	}
}

func (r *UserPostgresRepository) FindByID(id uuid.UUID) (*domain.User, error) {
	query := "SELECT id, username, password, role, status, created_at, updated_at FROM users WHERE id = $1"

	result := make(chan *domain.User)
	error := make(chan error)

	go func() {
		user := domain.User{}
		err := r.DB.QueryRow(query, id).Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.Role,
			&user.Status,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			error <- err
		}
		result <- &user
	}()

	select {
	case res := <-result:
		return res, nil
	case err := <-error:
		return nil, err
	}
}

func (r *UserPostgresRepository) Save(payload *domain.User) error {
	query := "INSERT INTO users (id, username, password, role, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, username, password, role, status, created_at, updated_at"

	hashed, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	payload.ID = uuid.New()
	payload.Password = string(hashed)
	payload.CreatedAt = time.Now()
	payload.UpdatedAt = time.Now()

	result := make(chan error)
	error := make(chan error)

	go func() {
		err := r.DB.QueryRow(query,
			payload.ID,
			payload.Username,
			payload.Password,
			payload.Role,
			payload.Status,
			payload.CreatedAt,
			payload.UpdatedAt,
		).Scan(
			&payload.ID,
			&payload.Username,
			&payload.Password,
			&payload.Role,
			&payload.Status,
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

func (r *UserPostgresRepository) Edit(payload *domain.User) error {
	query := "UPDATE users SET username = $1, password = $2, role = $3, status = $4, updated_at = $5 WHERE id = $6 RETURNING id, username, password, role, status, created_at, updated_at"

	if payload.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		payload.Password = string(hashed)
	}

	payload.UpdatedAt = time.Now()

	result := make(chan error)
	error := make(chan error)

	go func() {
		err := r.DB.QueryRow(query,
			payload.Username,
			payload.Password,
			payload.Role,
			payload.Status,
			payload.UpdatedAt,
			payload.ID,
		).Scan(
			&payload.ID,
			&payload.Username,
			&payload.Password,
			&payload.Role,
			&payload.Status,
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

func (r *UserPostgresRepository) EditPartial(payload *domain.User) error {
	query := "UPDATE users SET "
	args := []interface{}{}
	clauses := []string{}

	if payload.Username != "" {
		clauses = append(clauses, "username = $"+strconv.Itoa(len(args)+1))
		args = append(args, payload.Username)
	}

	if payload.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		clauses = append(clauses, "password = $"+strconv.Itoa(len(args)+1))
		args = append(args, string(hashed))
	}

	if payload.Role != "" {
		clauses = append(clauses, "role = $"+strconv.Itoa(len(args)+1))
		args = append(args, payload.Role)
	}

	if payload.Status != "" {
		clauses = append(clauses, "status = $"+strconv.Itoa(len(args)+1))
		args = append(args, payload.Status)
	}

	payload.UpdatedAt = time.Now()

	clauses = append(clauses, "updated_at = $"+strconv.Itoa(len(args)+1))
	args = append(args, payload.UpdatedAt)

	if len(clauses) == 0 {
		return nil
	}

	query += strings.Join(clauses, ", ") + " WHERE id = $" + strconv.Itoa(len(args)+1) + " RETURNING id, title, caption, created_at, updated_at"
	args = append(args, payload.ID)

	result := make(chan error)
	error := make(chan error)

	go func() {
		err := r.DB.QueryRow(query, args...).Scan(
			&payload.ID,
			&payload.Username,
			&payload.Password,
			&payload.Role,
			&payload.Status,
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

func (r *UserPostgresRepository) Remove(payload *domain.User) error {
	query := "DELETE FROM users WHERE id = $1"

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
