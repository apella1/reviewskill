-- +goose Up
CREATE TABLE
    flashcards (
        id UUID PRIMARY KEY,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        title TEXT NOT NULL,
        body TEXT NOT NULL,
        tags TEXT ARRAY, -- TEXT[] is an alternative
        last_reviewed_at TIMESTAMP,
        review_count INTEGER DEFAULT 0,
        correct_count INTEGER DEFAULT 0,
        difficulty_level INTEGER,
        user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE
    );

-- +goose Down
DROP TABLE flashcards;