package http

import (
	"encoding/json"
	"golang-solid-clean-architecture/entities"
	"golang-solid-clean-architecture/usecases"
	"net/http"
	"strconv"
)

type UserHandler struct {
	usecase *usecases.UserUsecase
}

func NewUserHandler(usecase *usecases.UserUsecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	json.NewDecoder(r.Body).Decode(&user)
	err := h.usecase.Register(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	user, err := h.usecase.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}
