package goMap

import (
	"fmt"
	"hw7garageproj/model"
	"slices"
	"strings"
	"sync"
)

var (
	storage *garage
	once    sync.Once
)

type garage struct {
	sync.Map
}

func GetGarage() *garage {
	once.Do(func() {
		storage = &garage{}
	})

	return storage
}

func (g *garage) AddUser(newUser *model.User) {
	g.Store(*newUser, []model.Vehicle{})
}

func (g *garage) UserCount() int {
	var count int
	g.Range(func(_, _ any) bool {
		count++
		return true
	})
	return count
}

func (g *garage) DeleteUser(id int) {
	g.Range(func(user, _ any) bool {
		if user.(model.User).Id == id {
			g.Delete(user)
			return false
		}
		return true
	})
}

func (g *garage) UserGet(id int) *model.User {
	var user model.User
	g.Range(func(u, _ any) bool {
		if u.(model.User).Id == id {
			user = u.(model.User)
			return false
		}
		return true
	})
	if user.Id == 0 && user.Name == "" {
		return nil
	} else {
		return &user
	}
}

func (g *garage) UpdateUser(newUser *model.User) {
	var oldUser model.User
	g.Range(func(u, _ any) bool {
		if u.(model.User).Id == newUser.Id {
			oldUser = u.(model.User)
			return false
		}
		return true
	})
	vehicles, _ := g.Load(oldUser)
	g.Delete(oldUser)
	g.Store(*newUser, vehicles)
}

func (g *garage) AddVehicle(newVehicle *model.Vehicle, toOwner *model.User) {
	vehicles, _ := g.Load(*toOwner)
	v := vehicles.([]model.Vehicle)
	g.Store(*toOwner, append(v, *newVehicle))
}

func (g *garage) VehicleById(id int) *model.Vehicle {
	var vehicle *model.Vehicle
	g.Range(func(_, vehicleSlice any) bool {
		for _, v := range vehicleSlice.([]model.Vehicle) {
			if v.Id == id {
				vehicle = &v
				return false
			}
		}
		return true
	})
	return vehicle
}

func (g *garage) DeleteVehicle(vehicle *model.Vehicle) {
	whose := g.whoseVehicle(vehicle)
	vehiclesAny, _ := g.Load(*whose)
	vehicleSlice := vehiclesAny.([]model.Vehicle)
	vehicleIdx := slices.Index(vehicleSlice, *vehicle)
	if vehicleIdx >= 0 {
		updatedSlice := append(vehicleSlice[:vehicleIdx], vehicleSlice[vehicleIdx+1:]...)
		g.Store(*whose, updatedSlice)
	}
}

func (g *garage) Vehicles(whose *model.User) []model.Vehicle {
	vehicle, _ := g.Load(*whose)
	return vehicle.([]model.Vehicle)
}

func (g *garage) UpdateVehicle(updatedVehicle *model.Vehicle) {
	whose := g.whoseVehicle(updatedVehicle)
	vAny, _ := g.Load(*whose)
	vehicleSlice := vAny.([]model.Vehicle)
	for i, vehicle := range vehicleSlice {
		if vehicle.Id == updatedVehicle.Id {
			vehicleSlice[i] = *updatedVehicle
		}
	}
}

func (g *garage) whoseVehicle(vehicle *model.Vehicle) *model.User {
	var owner model.User
	g.Range(func(user, vehicles any) bool {
		for _, v := range vehicles.([]model.Vehicle) {
			if v.Id == vehicle.Id {
				owner = user.(model.User)
				return false
			}
		}
		return true
	})
	return &owner
}

func (g *garage) String() string {
	var sb strings.Builder

	g.Range(func(u, v any) bool {
		user := u.(model.User)
		vehicles := v.([]model.Vehicle)
		sb.WriteString(fmt.Sprintf("%s with %d id has cars:\n", user.Name, user.Id))
		for _, vehicle := range vehicles {
			sb.WriteString(
				fmt.Sprintf("\tId:%d Name: %s Model:%s Brand: %s\n",
					vehicle.Id,
					vehicle.Name,
					vehicle.Model,
					vehicle.Brand,
				),
			)
		}
		return true
	})

	return sb.String()
}
