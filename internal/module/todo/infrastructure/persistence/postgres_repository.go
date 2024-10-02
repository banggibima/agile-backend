package persistence

import (
	"database/sql"
	"time"

	"github.com/banggibima/backend-agile/internal/module/todo/domain"
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
		query += " ORDER BY " + sort + " " + order
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
	query := "INSERT INTO todos (id, title, caption, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"

	payload.ID = uuid.New()
	payload.CreatedAt = time.Now()
	payload.UpdatedAt = time.Now()

	result := make(chan error)
	error := make(chan error)

	go func() {
		_, err := r.DB.Exec(
			query, payload.ID,
			payload.Title,
			payload.Caption,
			payload.CreatedAt,
			payload.UpdatedAt,
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
	query := "UPDATE todos SET title = $1, caption = $2, created_at = $3, updated_at = $4 WHERE id = $5"

	payload.UpdatedAt = time.Now()

	result := make(chan error)
	error := make(chan error)

	go func() {
		_, err := r.DB.Exec(query,
			payload.Title,
			payload.Caption,
			payload.CreatedAt,
			payload.UpdatedAt,
			payload.ID,
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
	query := "UPDATE todos SET title = $1, caption = $2, created_at = $3, updated_at = $4 WHERE id = $5"

	payload.UpdatedAt = time.Now()

	result := make(chan error)
	error := make(chan error)

	go func() {
		_, err := r.DB.Exec(query,
			payload.Title,
			payload.Caption,
			payload.CreatedAt,
			payload.UpdatedAt,
			payload.ID,
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
