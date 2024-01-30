package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reviewskill/config"
	"reviewskill/internal/database"
	"reviewskill/utils"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.RespondWithError(w, 500, fmt.Sprintf("Couldn't hash password: %v", err))
		return
	}
	user, err := h.Cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Email:     params.Email,
		Password:  string(hashedPassword),
	})
	if err != nil {
		utils.RespondWithError(w, 400, "User with the email exists!")
		return
	}
	utils.RespondWithJSON(w, 200, utils.DatabaseUserToUser(user))
}

func (h *Handler) UploadProfileImage(w http.ResponseWriter, r *http.Request, user database.User) {
	userIDStr := chi.URLParam(r, "id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	if user.ID != userID {
		http.Error(w, "User ID does not match authenticated user", http.StatusUnauthorized)
		return
	}
	err = r.ParseMultipartForm(1000)
	if err != nil {
		http.Error(w, "Maximum image size is 1MB", http.StatusBadRequest)
		return
	}
	file, _, err := r.FormFile("profile_image")
	if err != nil {
		http.Error(w, "An error occurred while retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "An error occurred while reading the file.", http.StatusBadRequest)
		return
	}
	_, err = h.Cfg.DB.SaveProfileImage(r.Context(), database.SaveProfileImageParams{
		ID:           userID,
		ProfileImage: fileBytes,
	})
	if err != nil {
		http.Error(w, "An error occurred while saving the image", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Profile image uploaded successfully"))
}
