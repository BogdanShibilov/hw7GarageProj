package vehicle

import (
	"hw7garageproj/handler/user"
	"hw7garageproj/model"
	"strings"
)

type Storage interface {
	AddVehicle(newVehicle *model.Vehicle, toOwner *model.User)
	DeleteVehicle(vehicle *model.Vehicle)
	Vehicles(whose *model.User) []model.Vehicle
	VehicleById(id int) *model.Vehicle
	UpdateVehicle(updatedVehicle *model.Vehicle)
}

func Add(vId int, vName, vBrand, vModel string,
	userId int, to Storage, userStorage user.Storage) error {
	if vId < 0 ||
		userId < 0 ||
		strings.Trim(vName, " ") == "" ||
		strings.Trim(vBrand, " ") == "" ||
		strings.Trim(vModel, " ") == "" {

		return InvalidInputErr
	}

	u, err := user.ById(userId, userStorage)
	if err != nil {
		return err
	}

	if vehicle := to.VehicleById(vId); vehicle != nil {
		return AlreadyExistsErr
	}

	newVehicle := &model.Vehicle{
		Id:    vId,
		Name:  vName,
		Model: vModel,
		Brand: vBrand,
	}

	to.AddVehicle(newVehicle, u)
	return nil
}

func AllBelongingTo(userId int, from Storage, userStorage user.Storage) ([]model.Vehicle, error) {
	if userId < 0 {
		return nil, InvalidInputErr
	}

	u, err := user.ById(userId, userStorage)
	if err != nil {
		return nil, err
	}

	return from.Vehicles(u), nil
}

func ById(id int, in Storage) (*model.Vehicle, error) {
	if id < 0 {
		return nil, InvalidInputErr
	}

	vehicle := in.VehicleById(id)
	if vehicle == nil {
		return nil, NotFoundErr
	}
	return vehicle, nil
}

func Update(vId int, vName, vBrand, vModel string, in Storage) error {
	if vId < 0 ||
		strings.Trim(vName, " ") == "" ||
		strings.Trim(vBrand, " ") == "" ||
		strings.Trim(vModel, " ") == "" {

		return InvalidInputErr
	}

	vehicle := in.VehicleById(vId)
	if vehicle == nil {
		return NotFoundErr
	}

	vehicle.Name = vName
	vehicle.Brand = vBrand
	vehicle.Model = vModel
	in.UpdateVehicle(vehicle)
	return nil
}

func Delete(vId int, in Storage) {
	if vId < 0 {
		return
	}

	vehicle := in.VehicleById(vId)
	in.DeleteVehicle(vehicle)
}
