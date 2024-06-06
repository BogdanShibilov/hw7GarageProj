package goMap

import (
	"hw7garageproj/internal/model"
	"slices"
)

type Garage map[model.User][]model.Vehicle

func (g Garage) AddUser(newUser *model.User) {
	g[*newUser] = make([]model.Vehicle, 0)
}

func (g Garage) DeleteUser(id int) {
	for user := range g {
		if user.Id == id {
			delete(g, user)
		}
	}
}

func (g Garage) User(id int) *model.User {
	for user := range g {
		if user.Id == id {
			return &user
		}
	}
	return nil
}

func (g Garage) UpdateUser(newUser *model.User) {
	oldUser := g.User(newUser.Id)
	g[*newUser] = g[*oldUser]
	delete(g, *oldUser)
}

func (g Garage) AddVehicle(newVehicle *model.Vehicle, toOwner *model.User) {
	g[*toOwner] = append(g[*toOwner], *newVehicle)
}

func (g Garage) VehicleById(id int) *model.Vehicle {
	for _, vehicleSlice := range g {
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
	vehicleSlice := g[*whose]
	vehicleIdx := slices.Index(vehicleSlice, *vehicle)
	if vehicleIdx >= 0 {
		g[*whose] = append(vehicleSlice[:vehicleIdx], vehicleSlice[vehicleIdx+1:]...)
	}
}

func (g Garage) Vehicles(whose *model.User) []model.Vehicle {
	return g[*whose]
}

func (g Garage) Vehicle(id int) *model.Vehicle {
	for _, vehicles := range g {
		for _, vehicle := range vehicles {
			if vehicle.Id == id {
				return &vehicle
			}
		}
	}
	return nil
}

func (g Garage) UpdateVehicle(updatedVehicle *model.Vehicle) {
	whose := g.whoseVehicle(updatedVehicle)
	vehicleSlice := g[*whose]
	for i, vehicle := range vehicleSlice {
		if vehicle.Id == updatedVehicle.Id {
			vehicleSlice[i] = *updatedVehicle
		}
	}
}

func (g Garage) whoseVehicle(vehicle *model.Vehicle) *model.User {
	for owner, vehicleSlice := range g {
		for _, v := range vehicleSlice {
			if v.Id == vehicle.Id {
				return &owner
			}
		}
	}
	return nil
}
