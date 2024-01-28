package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reviewskill/config"
	"reviewskill/internal/database"
	"reviewskill/utils"
	"time"

	"github.com/google/uuid"
)

type Handler struct {
	Cfg *config.ApiConfig
}

func (h *Handler) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	user, err := h.Cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Email:     params.Email,
		Password:  params.Password,
	})
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't create user %v", err))
		return
	}
	utils.RespondWithJSON(w, 200, user)
}