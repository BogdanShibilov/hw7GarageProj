package vehicle

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

func (h *Handler) AddVehicle(vId int, vName, vBrand, vModel string,
	userId int) error {
	var user *model.User
	if user = h.Storage.User(userId); user == nil {
		return errors.New("user not found")
	}

	if vehicle := h.Storage.VehicleById(vId); vehicle != nil {
		return errors.New("vehicle with such id already exists")
	}

	newVehicle := &model.Vehicle{
		Id:    vId,
		Name:  vName,
		Model: vModel,
		Brand: vBrand,
	}

	h.Storage.AddVehicle(newVehicle, user)
	return nil
}

func (h *Handler) Vehicles(userId int) ([]model.Vehicle, error) {
	var user *model.User
	if user = h.Storage.User(userId); user == nil {
		return nil, errors.New("user not found")
	}

	return h.Storage.Vehicles(user), nil
}

func (h *Handler) VehicleById(id int) (*model.Vehicle, error) {
	vehicle := h.Storage.VehicleById(id)
	if vehicle == nil {
		return nil, errors.New("vehicle not found")
	}
	return vehicle, nil
}

func (h *Handler) UpdateVehicle(vId int, vName, vBrand, vModel string) error {
	vehicle := h.Storage.Vehicle(vId)
	if vehicle == nil {
		return errors.New("vehicle with such id does not exist")
	}
	vehicle.Name = vName
	vehicle.Brand = vBrand
	vehicle.Model = vModel
	h.Storage.UpdateVehicle(vehicle)
	return nil
}

func (h *Handler) DeleteVehicle(vId int) {
	vehicle := h.Storage.VehicleById(vId)
	h.Storage.DeleteVehicle(vehicle)
}
