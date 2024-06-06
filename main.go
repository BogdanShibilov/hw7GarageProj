package main

import (
	"fmt"
	"hw7garageproj/handler/user"
	"hw7garageproj/handler/vehicle"
	"hw7garageproj/storage/goMap"
)

func main() {
	garage := goMap.NewGarage(10)

	_ = user.Add(0, "Jack", garage)
	_ = user.Add(1, "Smith", garage)
	_ = user.Add(2, "Maria", garage)

	err := user.Add(1, "Smith", garage)
	if err != nil {
		fmt.Printf("Error due to Smith %v\n\n", err)
	}

	fmt.Printf("Map storage before user update:\n%v\n\n", garage)
	_ = user.Update(1, "Stephan", garage)
	fmt.Printf("Map storage after user update:\n%v\n\n", garage)

	user.Delete(0, garage)
	fmt.Printf("Map storage after user delete:\n%v\n\n", garage)

	_ = vehicle.Add(0, "name0", "brand0", "model0", 2, garage, garage)
	_ = vehicle.Add(1, "name1", "brand1", "model1", 2, garage, garage)
	_ = vehicle.Add(2, "name2", "brand2", "model2", 1, garage, garage)
	fmt.Printf("Map storage after adding vehicles:\n%v\n\n", garage)

	if err = vehicle.Add(
		2, "name2", "brand2", "model2",
		1,
		garage, garage); err != nil {
		fmt.Printf("Error due adding vehicle with taken id: %v\n\n", err)
	}
	vehicles, _ := vehicle.AllBelongingTo(1, garage, garage)
	fmt.Printf("after adding vehicle with existing id:\n%v\n\n", vehicles)

	_ = vehicle.Update(2, "updatedName2", "brand2", "model2", garage)
	v, err := vehicle.AllBelongingTo(1, garage, garage)
	fmt.Printf("Updated vehicle of user 1: %v \n\n", v)

	v, err = vehicle.AllBelongingTo(2, garage, garage)
	fmt.Printf("vehicles of user 2 before delete: %v\n", v)
	vehicle.Delete(0, garage)
	v, err = vehicle.AllBelongingTo(2, garage, garage)
	fmt.Printf("vehicles of user 2 after delete: %v\n\n", v)
}
