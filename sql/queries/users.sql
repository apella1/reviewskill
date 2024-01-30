-- name: CreateUser :one
INSERT INTO
    users (
        id,
        created_at,
        updated_at,
        first_name,
        last_name,
        email,
        password
    )
VALUES ($1,$2,$3,$4,$5,$6,$7)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: SaveProfileImage :one
UPDATE users
SET profile_image = $1
WHERE id = $2
RETURNING *;