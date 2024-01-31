// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: flashcards.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const createFlashcard = `-- name: CreateFlashcard :one
INSERT INTO
    flashcards (
        id,
        created_at,
        updated_at,
        updated_at,
        title,
        body,
        tags,
        last_reviewed_at,
        review_count,
        correct_count,
        difficulty_level,
        user_id
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING id, created_at, updated_at, title, body, tags, last_reviewed_at, review_count, correct_count, difficulty_level, user_id
`

type CreateFlashcardParams struct {
	ID              uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time
	UpdatedAt_2     time.Time
	Title           string
	Body            string
	Tags            []string
	LastReviewedAt  sql.NullTime
	ReviewCount     sql.NullInt32
	CorrectCount    sql.NullInt32
	DifficultyLevel sql.NullInt32
	UserID          uuid.UUID
}

func (q *Queries) CreateFlashcard(ctx context.Context, arg CreateFlashcardParams) (Flashcard, error) {
	row := q.db.QueryRowContext(ctx, createFlashcard,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UpdatedAt_2,
		arg.Title,
		arg.Body,
		pq.Array(arg.Tags),
		arg.LastReviewedAt,
		arg.ReviewCount,
		arg.CorrectCount,
		arg.DifficultyLevel,
		arg.UserID,
	)
	var i Flashcard
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Body,
		pq.Array(&i.Tags),
		&i.LastReviewedAt,
		&i.ReviewCount,
		&i.CorrectCount,
		&i.DifficultyLevel,
		&i.UserID,
	)
	return i, err
}
