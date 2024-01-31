-- name: CreateFlashcard :one
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
RETURNING *;