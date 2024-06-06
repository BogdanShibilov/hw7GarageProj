package main

import (
	"fmt"
	"hw7garageproj/internal/handler/user"
	"hw7garageproj/internal/handler/vehicle"
	"hw7garageproj/internal/storage/goMap"
)

func main() {
	garage := goMap.Garage{}
	userHandler := user.NewHandler(garage)
	vehicleHandler := vehicle.NewHandler(garage)

	_ = userHandler.AddUser(0, "Jack")
	_ = userHandler.AddUser(1, "Smith")
	_ = userHandler.AddUser(2, "Maria")

	err := userHandler.AddUser(1, "Smith")
	if err != nil {
		fmt.Printf("Error due to Smith %v\n\n", err)
	}

	fmt.Printf("Map storage before user update:\n%v\n\n", userHandler.Storage)
	_ = userHandler.UpdateUser(1, "Stephan")
	fmt.Printf("Map storage after user update:\n%v\n\n", userHandler.Storage)

	userHandler.DeleteUser(0)
	fmt.Printf("Map storage after user delete:\n%v\n\n", userHandler.Storage)

	_ = vehicleHandler.AddVehicle(0, "name0", "brand0", "model0", 2)
	_ = vehicleHandler.AddVehicle(1, "name1", "brand1", "model1", 2)
	_ = vehicleHandler.AddVehicle(2, "name2", "brand2", "model2", 1)
	fmt.Printf("Map storage after adding vehicles:\n%v\n\n", userHandler.Storage)

	if err = vehicleHandler.AddVehicle(2, "name2", "brand2", "model2", 1); err != nil {
		fmt.Printf("Error due adding vehicle with taken id: %v\n\n", err)
	}
	vehicles, _ := vehicleHandler.Vehicles(1)
	fmt.Printf("after adding vehicle with existing id:\n%v\n\n", vehicles)

	_ = vehicleHandler.UpdateVehicle(2, "updatedName2", "brand2", "model2")
	v, err := vehicleHandler.Vehicles(1)
	fmt.Printf("Updated vehicle of user 1: %v \n\n", v)

	v, err = vehicleHandler.Vehicles(2)
	fmt.Printf("vehicles of user 2 before delete: %v\n", v)
	vehicleHandler.DeleteVehicle(0)
	v, err = vehicleHandler.Vehicles(2)
	fmt.Printf("vehicles of user 2 after delete: %v\n\n", v)
}
