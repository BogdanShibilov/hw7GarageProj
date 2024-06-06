package goMap

import (
	"hw7garageproj/model"
	"slices"
)

const (
	_defaultGarageCapacity = 10
)

type Garage struct {
	m map[model.User][]model.Vehicle
}

func NewGarage(capacity int) *Garage {
	if capacity <= 0 {
		capacity = _defaultGarageCapacity
	}

	return &Garage{
		m: make(map[model.User][]model.Vehicle, capacity),
	}
}

func (g Garage) AddUser(newUser *model.User) {
	g.m[*newUser] = make([]model.Vehicle, 0)
}

func (g Garage) UserCount() int {
	return len(g.m)
}

func (g Garage) DeleteUser(id int) {
	for user := range g.m {
		if user.Id == id {
			delete(g.m, user)
		}
	}
}

func (g Garage) User(id int) *model.User {
	for user := range g.m {
		if user.Id == id {
			return &user
		}
	}
	return nil
}

func (g Garage) UpdateUser(newUser *model.User) {
	oldUser := g.User(newUser.Id)
	g.m[*newUser] = g.m[*oldUser]
	delete(g.m, *oldUser)
}

func (g Garage) AddVehicle(newVehicle *model.Vehicle, toOwner *model.User) {
	g.m[*toOwner] = append(g.m[*toOwner], *newVehicle)
}

func (g Garage) VehicleById(id int) *model.Vehicle {
	for _, vehicleSlice := range g.m {
		for _, vehicle := range vehicleSlice {
			if vehicle.Id == id {
				return &vehicle
			}
		}
	}
	return nil
}

func (g Garage) DeleteVehicle(vehicle *model.Vehicle) {
	whose := g.whoseVehicle(vehicle)
	vehicleSlice := g.m[*whose]
	vehicleIdx := slices.Index(vehicleSlice, *vehicle)
	if vehicleIdx >= 0 {
		g.m[*whose] = append(vehicleSlice[:vehicleIdx], vehicleSlice[vehicleIdx+1:]...)
	}
}

func (g Garage) Vehicles(whose *model.User) []model.Vehicle {
	return g.m[*whose]
}

func (g Garage) UpdateVehicle(updatedVehicle *model.Vehicle) {
	whose := g.whoseVehicle(updatedVehicle)
	vehicleSlice := g.m[*whose]
	for i, vehicle := range vehicleSlice {
		if vehicle.Id == updatedVehicle.Id {
			vehicleSlice[i] = *updatedVehicle
		}
	}
}

func (g Garage) whoseVehicle(vehicle *model.Vehicle) *model.User {
	for owner, vehicleSlice := range g.m {
		for _, v := range vehicleSlice {
			if v.Id == vehicle.Id {
				return &owner
			}
		}
	}
	return nil
}
