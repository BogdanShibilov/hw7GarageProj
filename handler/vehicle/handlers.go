package vehicle

import (
	"hw7garageproj/handler/user"
	"hw7garageproj/model"
	"hw7garageproj/storage/goMap"
	"strings"
)

var storage = goMap.GetGarage()

func Add(vId int, vName, vBrand, vModel string, userId int) error {
	if err := validateVehicleParameters(vId, vName, vBrand, vModel); err != nil {
		return err
	}
	if userId < 0 {
		return user.InvalidIdErr
	}

	u, err := user.ById(userId)
	if err != nil {
		return err
	}

	if vehicle := storage.VehicleById(vId); vehicle != nil {
		return AlreadyExistsErr
	}

	newVehicle := &model.Vehicle{
		Id:    vId,
		Name:  vName,
		Model: vModel,
		Brand: vBrand,
	}

	storage.AddVehicle(newVehicle, u)
	return nil
}

func AllBelongingTo(userId int) ([]model.Vehicle, error) {
	if userId < 0 {
		return nil, user.InvalidIdErr
	}

	u, err := user.ById(userId)
	if err != nil {
		return nil, err
	}

	return storage.Vehicles(u), nil
}

func ById(id int) (*model.Vehicle, error) {
	if id < 0 {
		return nil, InvalidIdErr
	}

	vehicle := storage.VehicleById(id)
	if vehicle == nil {
		return nil, NotFoundErr
	}
	return vehicle, nil
}

func Update(vId int, vName, vBrand, vModel string) error {
	if err := validateVehicleParameters(vId, vName, vBrand, vModel); err != nil {
		return err
	}

	vehicle := storage.VehicleById(vId)
	if vehicle == nil {
		return NotFoundErr
	}

	vehicle.Name = vName
	vehicle.Brand = vBrand
	vehicle.Model = vModel
	storage.UpdateVehicle(vehicle)
	return nil
}

func Delete(vId int) {
	if vId < 0 {
		return
	}

	vehicle := storage.VehicleById(vId)
	storage.DeleteVehicle(vehicle)
}

func validateVehicleParameters(vId int, vName, vBrand, vModel string) error {
	if vId < 0 {
		return InvalidIdErr
	}
	if strings.Trim(vName, " ") == "" {
		return InvalidNameErr
	}
	if strings.Trim(vBrand, " ") == "" {
		return InvalidBrandErr
	}
	if strings.Trim(vModel, " ") == "" {
		return InvalidModelErr
	}
	return nil
}
