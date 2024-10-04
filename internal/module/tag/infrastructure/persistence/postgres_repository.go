package persistence

import (
	"database/sql"
	"strconv"
	"strings"
	"time"

	"github.com/banggibima/agile-backend/internal/module/tag/domain"
	"github.com/google/uuid"
)

type TagPostgresRepository struct {
	DB *sql.DB
}

func NewTagPostgresRepository(db *sql.DB) *TagPostgresRepository {
	return &TagPostgresRepository{
		DB: db,
	}
}

func (r *TagPostgresRepository) Count() (int, error) {
	total := 0
	query := "SELECT COUNT(*) FROM tags"

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

func (r *TagPostgresRepository) Find(offset, limit int, sort, order string) ([]*domain.Tag, error) {
	query := "SELECT id, name, created_at, updated_at FROM tags"

	if sort != "" && order != "" {
		query += " ORDER BY " + sort + " " + strings.ToUpper(order)
	} else {
		query += " ORDER BY created_at DESC"
	}

	if limit > 0 && offset >= 0 {
		query += " LIMIT $1 OFFSET $2"

		result := make(chan []*domain.Tag)
		error := make(chan error)

		go func() {
			rows, err := r.DB.Query(query, limit, offset)
			if err != nil {
				error <- err
			}
			defer rows.Close()

			tags := []*domain.Tag{}
			for rows.Next() {
				tag := domain.Tag{}
				if err := rows.Scan(
					&tag.ID,
					&tag.Name,
					&tag.CreatedAt,
					&tag.UpdatedAt,
				); err != nil {
					error <- err
				}
				tags = append(tags, &tag)
			}

			if err := rows.Err(); err != nil {
				error <- err
			}

			result <- tags
		}()

		select {
		case res := <-result:
			return res, nil
		case err := <-error:
			return nil, err
		}
	} else {

		result := make(chan []*domain.Tag)
		error := make(chan error)

		go func() {
			rows, err := r.DB.Query(query)
			if err != nil {
				error <- err
			}
			defer rows.Close()

			tags := []*domain.Tag{}
			for rows.Next() {
				tag := domain.Tag{}
				if err := rows.Scan(
					&tag.ID,
					&tag.Name,
					&tag.CreatedAt,
					&tag.UpdatedAt,
				); err != nil {
					error <- err
				}
				tags = append(tags, &tag)
			}

			if err := rows.Err(); err != nil {
				error <- err
			}

			result <- tags
		}()

		select {
		case res := <-result:
			return res, nil
		case err := <-error:
			return nil, err
		}
	}
}

func (r *TagPostgresRepository) FindByID(id uuid.UUID) (*domain.Tag, error) {
	query := "SELECT id, name, created_at, updated_at FROM tags WHERE id = $1"

	result := make(chan *domain.Tag)
	error := make(chan error)

	go func() {
		tag := domain.Tag{}
		err := r.DB.QueryRow(query, id).Scan(
			&tag.ID,
			&tag.Name,
			&tag.CreatedAt,
			&tag.UpdatedAt,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				result <- nil
			}
			error <- err
		}
		result <- &tag
	}()

	select {
	case res := <-result:
		return res, nil
	case err := <-error:
		return nil, err
	}
}

func (r *TagPostgresRepository) Save(payload *domain.Tag) error {
	query := "INSERT INTO tags (id, name, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id, name, created_at, updated_at"

	payload.ID = uuid.New()
	payload.CreatedAt = time.Now()
	payload.UpdatedAt = time.Now()

	result := make(chan error)
	error := make(chan error)

	go func() {
		err := r.DB.QueryRow(query,
			payload.ID,
			payload.Name,
			payload.CreatedAt,
			payload.UpdatedAt,
		).Scan(
			&payload.ID,
			&payload.Name,
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

func (r *TagPostgresRepository) Edit(payload *domain.Tag) error {
	query := "UPDATE tags SET name = $1, updated_at = $2 WHERE id = $3 RETURNING id, name, created_at, updated_at"

	payload.UpdatedAt = time.Now()

	result := make(chan error)
	error := make(chan error)

	go func() {
		err := r.DB.QueryRow(query,
			payload.Name,
			payload.UpdatedAt,
			payload.ID,
		).Scan(
			&payload.ID,
			&payload.Name,
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

func (r *TagPostgresRepository) EditPartial(payload *domain.Tag) error {
	query := "UPDATE tags SET "
	args := []interface{}{}
	clauses := []string{}

	if payload.Name != "" {
		clauses = append(clauses, "name = $"+strconv.Itoa(len(args)+1))
		args = append(args, payload.Name)
	}

	payload.UpdatedAt = time.Now()

	clauses = append(clauses, "updated_at = $"+strconv.Itoa(len(args)+1))
	args = append(args, payload.UpdatedAt)

	if len(clauses) == 0 {
		return nil
	}

	query += strings.Join(clauses, ", ") + " WHERE id = $" + strconv.Itoa(len(args)+1) + " RETURNING id, name, created_at, updated_at"
	args = append(args, payload.ID)

	result := make(chan error)
	error := make(chan error)

	go func() {
		err := r.DB.QueryRow(query, args...).Scan(
			&payload.ID,
			&payload.Name,
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

func (r *TagPostgresRepository) Remove(payload *domain.Tag) error {
	query := "DELETE FROM tags WHERE id = $1"

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
