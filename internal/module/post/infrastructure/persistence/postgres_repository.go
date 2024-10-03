package persistence

import (
	"database/sql"
	"strconv"
	"strings"
	"time"

	"github.com/banggibima/agile-backend/internal/module/post/domain"
	"github.com/google/uuid"
)

type PostPostgresRepository struct {
	DB *sql.DB
}

func NewPostPostgresRepository(db *sql.DB) *PostPostgresRepository {
	return &PostPostgresRepository{
		DB: db,
	}
}

func (r *PostPostgresRepository) Count() (int, error) {
	total := 0
	query := "SELECT COUNT(*) FROM posts"

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

func (r *PostPostgresRepository) Find(offset, limit int, sort, order string) ([]*domain.Post, error) {
	query := "SELECT id, title, content, user_id, created_at, updated_at FROM posts"

	if sort != "" && order != "" {
		query += " ORDER BY " + sort + " " + strings.ToUpper(order)
	} else {
		query += " ORDER BY created_at DESC"
	}

	if limit > 0 && offset >= 0 {
		query += " LIMIT $1 OFFSET $2"

		result := make(chan []*domain.Post)
		error := make(chan error)

		go func() {
			rows, err := r.DB.Query(query, limit, offset)
			if err != nil {
				error <- err
			}
			defer rows.Close()

			posts := []*domain.Post{}
			for rows.Next() {
				post := domain.Post{}
				if err := rows.Scan(
					&post.ID,
					&post.Title,
					&post.Content,
					&post.UserID,
					&post.CreatedAt,
					&post.UpdatedAt,
				); err != nil {
					error <- err
				}
				posts = append(posts, &post)
			}

			if err := rows.Err(); err != nil {
				error <- err
			}

			result <- posts
		}()

		select {
		case res := <-result:
			return res, nil
		case err := <-error:
			return nil, err
		}
	} else {

		result := make(chan []*domain.Post)
		error := make(chan error)

		go func() {
			rows, err := r.DB.Query(query)
			if err != nil {
				error <- err
			}
			defer rows.Close()

			posts := []*domain.Post{}
			for rows.Next() {
				post := domain.Post{}
				if err := rows.Scan(
					&post.ID,
					&post.Title,
					&post.Content,
					&post.UserID,
					&post.CreatedAt,
					&post.UpdatedAt,
				); err != nil {
					error <- err
				}
				posts = append(posts, &post)
			}

			if err := rows.Err(); err != nil {
				error <- err
			}

			result <- posts
		}()

		select {
		case res := <-result:
			return res, nil
		case err := <-error:
			return nil, err
		}
	}
}

func (r *PostPostgresRepository) FindByID(id uuid.UUID) (*domain.Post, error) {
	query := "SELECT id, title, content, user_id, created_at, updated_at FROM posts WHERE id = $1"

	result := make(chan *domain.Post)
	error := make(chan error)

	go func() {
		post := domain.Post{}
		err := r.DB.QueryRow(query, id).Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.UserID,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			error <- err
		}
		result <- &post
	}()

	select {
	case res := <-result:
		return res, nil
	case err := <-error:
		return nil, err
	}
}

func (r *PostPostgresRepository) Save(payload *domain.Post) error {
	query := "INSERT INTO posts (id, title, content, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, title, content, user_id, created_at, updated_at"

	payload.ID = uuid.New()
	payload.CreatedAt = time.Now()
	payload.UpdatedAt = time.Now()

	result := make(chan error)
	error := make(chan error)

	go func() {
		err := r.DB.QueryRow(query,
			payload.ID,
			payload.Title,
			payload.Content,
			payload.UserID,
			payload.CreatedAt,
			payload.UpdatedAt,
		).Scan(
			&payload.ID,
			&payload.Title,
			&payload.Content,
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

func (r *PostPostgresRepository) Edit(payload *domain.Post) error {
	query := "UPDATE posts SET title = $1, content = $2, user_id = $3, updated_at = $4 WHERE id = $5 RETURNING id, title, content, user_id, created_at, updated_at"

	payload.UpdatedAt = time.Now()

	result := make(chan error)
	error := make(chan error)

	go func() {
		err := r.DB.QueryRow(query,
			payload.Title,
			payload.Content,
			payload.UserID,
			payload.UpdatedAt,
			payload.ID,
		).Scan(
			&payload.ID,
			&payload.Title,
			&payload.Content,
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

func (r *PostPostgresRepository) EditPartial(payload *domain.Post) error {
	query := "UPDATE posts SET "
	args := []interface{}{}
	clauses := []string{}

	if payload.Title != "" {
		clauses = append(clauses, "title = $"+strconv.Itoa(len(args)+1))
		args = append(args, payload.Title)
	}

	if payload.Content != "" {
		clauses = append(clauses, "content = $"+strconv.Itoa(len(args)+1))
		args = append(args, payload.Content)
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

	query += strings.Join(clauses, ", ") + " WHERE id = $" + strconv.Itoa(len(args)+1) + " RETURNING id, title, content, user_id, created_at, updated_at"
	args = append(args, payload.ID)

	result := make(chan error)
	error := make(chan error)

	go func() {
		err := r.DB.QueryRow(query, args...).Scan(
			&payload.ID,
			&payload.Title,
			&payload.Content,
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

func (r *PostPostgresRepository) Remove(payload *domain.Post) error {
	query := "DELETE FROM posts WHERE id = $1"

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
