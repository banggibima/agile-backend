package persistence

import (
	"database/sql"
	"strconv"
	"strings"
	"time"

	"github.com/banggibima/agile-backend/internal/module/todo/domain"
	"github.com/google/uuid"
)

type TodoPostgresRepository struct {
	DB *sql.DB
}

func NewTodoPostgresRepository(db *sql.DB) *TodoPostgresRepository {
	return &TodoPostgresRepository{
		DB: db,
	}
}

func (r *TodoPostgresRepository) Count() (int, error) {
	total := 0
	query := "SELECT COUNT(*) FROM todos"

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

func (r *TodoPostgresRepository) Find(offset, limit int, sort, order string) ([]*domain.Todo, error) {
	query := "SELECT id, title, caption, created_at, updated_at FROM todos"

	if sort != "" && order != "" {
		query += " ORDER BY " + sort + " " + strings.ToUpper(order)
	} else {
		query += " ORDER BY created_at DESC"
	}

	if limit > 0 && offset >= 0 {
		query += " LIMIT $1 OFFSET $2"

		result := make(chan []*domain.Todo)
		error := make(chan error)

		go func() {
			rows, err := r.DB.Query(query, limit, offset)
			if err != nil {
				error <- err
			}
			defer rows.Close()

			todos := []*domain.Todo{}
			for rows.Next() {
				todo := domain.Todo{}
				if err := rows.Scan(
					&todo.ID,
					&todo.Title,
					&todo.Caption,
					&todo.CreatedAt,
					&todo.UpdatedAt,
				); err != nil {
					error <- err
				}
				todos = append(todos, &todo)
			}

			if err := rows.Err(); err != nil {
				error <- err
			}

			result <- todos
		}()

		select {
		case res := <-result:
			return res, nil
		case err := <-error:
			return nil, err
		}
	} else {

		result := make(chan []*domain.Todo)
		error := make(chan error)

		go func() {
			rows, err := r.DB.Query(query)
			if err != nil {
				error <- err
			}
			defer rows.Close()

			todos := []*domain.Todo{}
			for rows.Next() {
				todo := domain.Todo{}
				if err := rows.Scan(
					&todo.ID,
					&todo.Title,
					&todo.Caption,
					&todo.CreatedAt,
					&todo.UpdatedAt,
				); err != nil {
					error <- err
				}
				todos = append(todos, &todo)
			}

			if err := rows.Err(); err != nil {
				error <- err
			}

			result <- todos
		}()

		select {
		case res := <-result:
			return res, nil
		case err := <-error:
			return nil, err
		}
	}
}

func (r *TodoPostgresRepository) FindByID(id uuid.UUID) (*domain.Todo, error) {
	query := "SELECT id, title, caption, created_at, updated_at FROM todos WHERE id = $1"

	result := make(chan *domain.Todo)
	error := make(chan error)

	go func() {
		todo := domain.Todo{}
		err := r.DB.QueryRow(query, id).Scan(
			&todo.ID,
			&todo.Title,
			&todo.Caption,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				result <- nil
			}
			error <- err
		}
		result <- &todo
	}()

	select {
	case res := <-result:
		return res, nil
	case err := <-error:
		return nil, err
	}
}

func (r *TodoPostgresRepository) Save(payload *domain.Todo) error {
	query := "INSERT INTO todos (id, title, caption, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, title, caption, created_at, updated_at"

	payload.ID = uuid.New()
	payload.CreatedAt = time.Now()
	payload.UpdatedAt = time.Now()

	result := make(chan error)
	error := make(chan error)

	go func() {
		err := r.DB.QueryRow(query,
			payload.ID,
			payload.Title,
			payload.Caption,
			payload.CreatedAt,
			payload.UpdatedAt,
		).Scan(
			&payload.ID,
			&payload.Title,
			&payload.Caption,
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

func (r *TodoPostgresRepository) Edit(payload *domain.Todo) error {
	query := "UPDATE todos SET title = $1, caption = $2, updated_at = $3 WHERE id = $4 RETURNING id, title, caption, created_at, updated_at"

	payload.UpdatedAt = time.Now()

	result := make(chan error)
	error := make(chan error)

	go func() {
		err := r.DB.QueryRow(query,
			payload.Title,
			payload.Caption,
			payload.UpdatedAt,
			payload.ID,
		).Scan(
			&payload.ID,
			&payload.Title,
			&payload.Caption,
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

func (r *TodoPostgresRepository) EditPartial(payload *domain.Todo) error {
	query := "UPDATE todos SET "
	args := []interface{}{}
	clauses := []string{}

	if payload.Title != "" {
		clauses = append(clauses, "title = $"+strconv.Itoa(len(args)+1))
		args = append(args, payload.Title)
	}

	if payload.Caption != "" {
		clauses = append(clauses, "caption = $"+strconv.Itoa(len(args)+1))
		args = append(args, payload.Caption)
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
			&payload.Title,
			&payload.Caption,
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

func (r *TodoPostgresRepository) Remove(payload *domain.Todo) error {
	query := "DELETE FROM todos WHERE id = $1"

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
