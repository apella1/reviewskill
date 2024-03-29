package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reviewskill/internal/database"
	"reviewskill/utils"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (h *Handler) HandlerCreateFlashcard(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Title string   `json:"title"`
		Body  string   `json:"body"`
		Tags  []string `json:"tags"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	flashcard, err := h.Cfg.DB.CreateFlashcard(r.Context(), database.CreateFlashcardParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Title:     params.Title,
		Body:      params.Body,
		Tags:      params.Tags,
		UserID:    user.ID,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusForbidden, "Error creating flashcard")
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, utils.DatabaseFlashcardToFlashcard(flashcard))
}

func (h *Handler) HandlerFetchUserFlashcards(w http.ResponseWriter, r *http.Request, user database.User) {
	flashcards, err := h.Cfg.DB.FetchUserFlashcards(r.Context(), user.ID)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't fetch user flashcards!: %v", err))
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, utils.DatabaseFlashcardsToFlashcards(flashcards))
}

func (h *Handler) HandlerDeleteFlashcard(w http.ResponseWriter, r *http.Request, user database.User) {
	flashcardIDstr := chi.URLParam(r, "id")
	flashcardID, err := uuid.Parse(flashcardIDstr)
	if err != nil {
		utils.RespondWithError(w, http.StatusForbidden, "Invalid flashcard ID")
	}
	err = h.Cfg.DB.DeleteFlashcard(r.Context(), database.DeleteFlashcardParams{
		ID:     flashcardID,
		UserID: user.ID,
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			utils.RespondWithError(w, http.StatusNotFound, "Flash card not found or already deleted!")
		} else {
			utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't delete flashcard: %v", err))
		}
	}
	utils.RespondWithJSON(w, http.StatusOK, "Deleted flashcard successfully")
}
