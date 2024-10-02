package persistence

import (
	"database/sql"
	"time"

	"github.com/banggibima/backend-agile/internal/module/user/domain"
	"github.com/google/uuid"
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

	err := r.DB.QueryRow(query).Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (r *UserPostgresRepository) Find(offset, limit int, sort, order string) ([]*domain.User, error) {
	query := "SELECT id, username, password, role, status, created_at, updated_at FROM users"

	if sort != "" && order != "" {
		query += " ORDER BY " + sort + " " + order
	} else {
		query += " ORDER BY created_at DESC"
	}

	if limit > 0 && offset >= 0 {
		query += " LIMIT $1 OFFSET $2"
		rows, err := r.DB.Query(query, limit, offset)
		if err != nil {
			return nil, err
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
				return nil, err
			}
			users = append(users, &user)
		}

		if err := rows.Err(); err != nil {
			return nil, err
		}

		return users, nil
	} else {
		rows, err := r.DB.Query(query)
		if err != nil {
			return nil, err
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
				return nil, err
			}
			users = append(users, &user)
		}

		if err := rows.Err(); err != nil {
			return nil, err
		}

		return users, nil
	}
}

func (r *UserPostgresRepository) FindByID(id uuid.UUID) (*domain.User, error) {
	query := "SELECT id, username, password, role, status, created_at, updated_at FROM users WHERE id = $1"

	row := r.DB.QueryRow(query, id)
	user := domain.User{}
	if err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserPostgresRepository) Save(payload *domain.User) error {
	query := "INSERT INTO users (id, username, password, role, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	payload.ID = uuid.New()
	payload.CreatedAt = time.Now()
	payload.UpdatedAt = time.Now()

	_, err := r.DB.Exec(
		query, payload.ID,
		payload.Username,
		payload.Password,
		payload.Role,
		payload.Status,
		payload.CreatedAt,
		payload.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserPostgresRepository) Edit(payload *domain.User) error {
	query := "UPDATE users SET username = $1, password = $2, role = $3, status = $4, created_at = $5, updated_at = $6 WHERE id = $7"

	payload.UpdatedAt = time.Now()

	_, err := r.DB.Exec(query,
		payload.Username,
		payload.Password,
		payload.Role,
		payload.Status,
		payload.CreatedAt,
		payload.UpdatedAt,
		payload.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserPostgresRepository) EditPartial(payload *domain.User) error {
	query := "UPDATE users SET username = $1, password = $2, role = $3, status = $4, created_at = $5, updated_at = $6 WHERE id = $7"

	payload.UpdatedAt = time.Now()

	_, err := r.DB.Exec(query,
		payload.Username,
		payload.Password,
		payload.Role,
		payload.Status,
		payload.CreatedAt,
		payload.UpdatedAt,
		payload.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserPostgresRepository) Remove(payload *domain.User) error {
	query := "DELETE FROM users WHERE id = $1"

	_, err := r.DB.Exec(query, payload.ID)
	if err != nil {
		return err
	}
	return nil
}
