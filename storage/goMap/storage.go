package goMap

import (
	"hw7garageproj/model"
	"slices"
	"sync"
)

const (
	_defaultGarageCapacity = 10
)

var (
	storage garage
	lock    = &sync.Mutex{}
)

type garage map[model.User][]model.Vehicle

func GetGarage() *garage {
	if storage == nil {
		lock.Lock()
		defer lock.Unlock()
		if storage == nil {
			storage = make(garage, _defaultGarageCapacity)
		}
	}

	return &storage
}

func (g garage) AddUser(newUser *model.User) {
	g[*newUser] = make([]model.Vehicle, 0)
}

func (g garage) UserCount() int {
	return len(g)
}

func (g garage) DeleteUser(id int) {
	for user := range g {
		if user.Id == id {
			delete(g, user)
		}
	}
}

func (g garage) User(id int) *model.User {
	for user := range g {
		if user.Id == id {
			return &user
		}
	}
	return nil
}

func (g garage) UpdateUser(newUser *model.User) {
	oldUser := g.User(newUser.Id)
	g[*newUser] = g[*oldUser]
	delete(g, *oldUser)
}

func (g garage) AddVehicle(newVehicle *model.Vehicle, toOwner *model.User) {
	g[*toOwner] = append(g[*toOwner], *newVehicle)
}

func (g garage) VehicleById(id int) *model.Vehicle {
	for _, vehicleSlice := range g {
		for _, vehicle := range vehicleSlice {
			if vehicle.Id == id {
				return &vehicle
			}
		}
	}
	return nil
}

func (g garage) DeleteVehicle(vehicle *model.Vehicle) {
	whose := g.whoseVehicle(vehicle)
	vehicleSlice := g[*whose]
	vehicleIdx := slices.Index(vehicleSlice, *vehicle)
	if vehicleIdx >= 0 {
		g[*whose] = append(vehicleSlice[:vehicleIdx], vehicleSlice[vehicleIdx+1:]...)
	}
}

func (g garage) Vehicles(whose *model.User) []model.Vehicle {
	return g[*whose]
}

func (g garage) UpdateVehicle(updatedVehicle *model.Vehicle) {
	whose := g.whoseVehicle(updatedVehicle)
	vehicleSlice := g[*whose]
	for i, vehicle := range vehicleSlice {
		if vehicle.Id == updatedVehicle.Id {
			vehicleSlice[i] = *updatedVehicle
		}
	}
}

func (g garage) whoseVehicle(vehicle *model.Vehicle) *model.User {
	for owner, vehicleSlice := range g {
		for _, v := range vehicleSlice {
			if v.Id == vehicle.Id {
				return &owner
			}
		}
	}
	return nil
}
