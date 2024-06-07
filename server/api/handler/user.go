package handler

import (
	"encoding/json"
	"net/http"
	"server/models"
	"server/service"
	"server/storage"
	"strconv"
)

type Handler struct {
	Storage *storage.User
	UserService *service.UserService
}

func NewHandler(s *storage.User, u *service.UserService) *Handler {
	return &Handler{Storage: s, UserService: u}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var body models.UserRequest

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.Storage.CreateUser(body)
	if err != nil {
		http.Error(w, "failed creating new user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	resp, err := h.Storage.GetAllUsers()
	if err != nil {
		http.Error(w, "failed getting all users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}
	resp, err := h.Storage.GetUserByID(intID)
	if err != nil {
		http.Error(w, "failed getting user by id", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "failed getting user by id", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	var body models.UserRequest

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.Storage.UpdateUserByID(intID, body)
	if err != nil {
		http.Error(w, "failed getting user by id", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	resp, err := h.Storage.DeleteUserByID(intID)
	if err != nil {
		http.Error(w, "failed getting user by id", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
