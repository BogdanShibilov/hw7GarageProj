package user

import (
	"errors"
	"hw7garageproj/internal/model"
	"hw7garageproj/internal/storage/goMap"
)

type Handler struct {
	Storage goMap.Garage
}

func NewHandler(storage goMap.Garage) *Handler {
	return &Handler{
		Storage: storage,
	}
}

func (h *Handler) AddUser(id int, name string) error {
	if user := h.Storage.User(id); user != nil {
		return errors.New("user already exists")
	}

	newUser := model.User{
		Id:   id,
		Name: name,
	}

	h.Storage.AddUser(&newUser)
	return nil
}

func (h *Handler) DeleteUser(id int) {
	h.Storage.DeleteUser(id)
}

func (h *Handler) User(id int) *model.User {
	return h.Storage.User(id)
}

func (h *Handler) UpdateUser(userId int, newName string) error {
	var user *model.User
	if user = h.Storage.User(userId); user == nil {
		return errors.New("user not found")
	}
	user.Name = newName
	h.Storage.UpdateUser(user)
	return nil
}
