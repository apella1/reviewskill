package utils

import (
	"database/sql"
	"reviewskill/internal/database"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	ProfileImage []byte    `json:"profile_image"`
}

func DatabaseUserToUser(dbUser database.User) User {
	return User{
		ID:           dbUser.ID,
		CreatedAt:    dbUser.CreatedAt,
		UpdatedAt:    dbUser.UpdatedAt,
		FirstName:    dbUser.FirstName,
		LastName:     dbUser.LastName,
		Email:        dbUser.Email,
		Password:     dbUser.Password,
		ProfileImage: dbUser.ProfileImage,
	}
}

type Flashcard struct {
	ID              uuid.UUID     `json:"id"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
	Title           string        `json:"title"`
	Body            string        `json:"body"`
	Tags            []string      `json:"tags"`
	LastReviewedAt  sql.NullTime  `json:"last_reviewed_at"`
	ReviewCount     sql.NullInt32 `json:"review_count"`
	CorrectCount    sql.NullInt32 `json:"correct_count"`
	DifficultyLevel sql.NullInt32 `json:"difficulty_level"`
	UserID          uuid.UUID     `json:"user_id"`
}

func DatabaseFlashcardToFlashcard(dbFlashcard database.Flashcard) Flashcard {
	return Flashcard{
		ID:              dbFlashcard.ID,
		CreatedAt:       dbFlashcard.CreatedAt,
		UpdatedAt:       dbFlashcard.UpdatedAt,
		Title:           dbFlashcard.Title,
		Body:            dbFlashcard.Body,
		Tags:            dbFlashcard.Tags,
		LastReviewedAt:  dbFlashcard.LastReviewedAt,
		ReviewCount:     dbFlashcard.ReviewCount,
		CorrectCount:    dbFlashcard.CorrectCount,
		DifficultyLevel: dbFlashcard.DifficultyLevel,
		UserID:          dbFlashcard.UserID,
	}
}

func DatabaseFlashcardsToFlashcards(dbFlashcards []database.Flashcard) []Flashcard {
	flashcards := []Flashcard{}
	for _, dbFlashcard := range dbFlashcards {
		flashcards = append(flashcards, Flashcard(dbFlashcard))
	}
	return flashcards
}
