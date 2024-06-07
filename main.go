package main

import (
	"fmt"
	"hw7garageproj/handler/user"
	"hw7garageproj/handler/vehicle"
	"hw7garageproj/storage/goMap"
	"time"
)

func main() {
	go func() {
		_ = user.Add(3, "ShortName")
	}()
	go func() {
		_ = user.Add(4, "LongName")
	}()
	time.Sleep(1 * time.Second)
	_ = user.Add(0, "Jack")
	_ = user.Add(1, "Smith")
	_ = user.Add(2, "Maria")

	err := user.Add(1, "Smith")
	if err != nil {
		fmt.Printf("Error due to Smith %v\n\n", err)
	}

	fmt.Printf("Map storage before user update:\n%v\n\n", goMap.GetGarage())
	_ = user.Update(1, "Stephan")
	fmt.Printf("Map storage after user update:\n%v\n\n", goMap.GetGarage())

	user.Delete(0)
	fmt.Printf("Map storage after user delete:\n%v\n\n", goMap.GetGarage())

	_ = vehicle.Add(0, "name0", "brand0", "model0", 2)
	_ = vehicle.Add(1, "name1", "brand1", "model1", 2)
	_ = vehicle.Add(2, "name2", "brand2", "model2", 1)
	fmt.Printf("Map storage after adding vehicles:\n%v\n\n", goMap.GetGarage())

	if err = vehicle.Add(
		2, "name2", "brand2", "model2",
		1); err != nil {
		fmt.Printf("Error due adding vehicle with taken id: %v\n\n", err)
	}
	vehicles, _ := vehicle.AllBelongingTo(1)
	fmt.Printf("after adding vehicle with existing id:\n%v\n\n", vehicles)

	_ = vehicle.Update(2, "updatedName2", "brand2", "model2")
	v, err := vehicle.AllBelongingTo(1)
	fmt.Printf("Updated vehicle of user 1: %v \n\n", v)

	v, err = vehicle.AllBelongingTo(2)
	fmt.Printf("vehicles of user 2 before delete: %v\n", v)
	vehicle.Delete(0)
	v, err = vehicle.AllBelongingTo(2)
	fmt.Printf("vehicles of user 2 after delete: %v\n\n", v)
}
