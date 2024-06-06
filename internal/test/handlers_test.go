package test

import (
	"github.com/stretchr/testify/require"
	"hw7garageproj/internal/handler/user"
	"hw7garageproj/internal/handler/vehicle"
	"hw7garageproj/internal/storage/goMap"
	"testing"
)

func TestList(t *testing.T) {
	t.Run("handlers", func(t *testing.T) {
		garage := goMap.Garage{}
		userHandler := user.NewHandler(garage)

		_ = userHandler.AddUser(0, "John")
		_ = userHandler.AddUser(1, "Jack")
		err := userHandler.AddUser(2, "Janny")

		require.NoError(t, err)
		require.Equal(t, 3, len(garage))

		userHandler.DeleteUser(0)
		require.Equal(t, 2, len(garage))

		userHandler.DeleteUser(5)
		require.Equal(t, 2, len(garage))

		err = userHandler.UpdateUser(2, "Johny")
		require.NoError(t, err)
		johnyUser := userHandler.User(2)
		require.NotNil(t, johnyUser)
		require.Equal(t, "Johny", johnyUser.Name)

		err = userHandler.AddUser(2, "Jennyfer")
		require.NotNil(t, err)
		require.Equal(t, 2, len(garage))

		err = userHandler.UpdateUser(999, "Jest")
		require.NotNil(t, err)

		vehicleHandler := vehicle.NewHandler(garage)

		err = vehicleHandler.AddVehicle(0, "vName1", "vBrand1", "vModel1", 1)
		require.NoError(t, err)

		vehicles, err := vehicleHandler.Vehicles(1)
		require.NoError(t, err)
		require.Equal(t, 1, len(vehicles))

		err = vehicleHandler.AddVehicle(0, "vName1", "vBrand1", "vModel1", 1)
		require.NotNil(t, err)

		vehicles, err = vehicleHandler.Vehicles(1)
		require.NoError(t, err)
		require.Equal(t, 1, len(vehicles))

		err = vehicleHandler.AddVehicle(1, "vName2", "vBrand2", "vModel2", 1)
		require.NoError(t, err)
		vehicles, err = vehicleHandler.Vehicles(1)
		require.NoError(t, err)
		require.Equal(t, 2, len(vehicles))

		v1, err := vehicleHandler.VehicleById(1)
		require.NoError(t, err)
		require.Equal(t, "vName2", v1.Name)

		err = vehicleHandler.UpdateVehicle(1, "updatedVName2", "vBrand2", "vModel2")
		require.NoError(t, err)

		updatedV1, err := vehicleHandler.VehicleById(1)
		require.NoError(t, err)
		require.Equal(t, "updatedVName2", updatedV1.Name)

		err = vehicleHandler.UpdateVehicle(999, "no", "no", "no")
		require.NotNil(t, err)

		vehicleHandler.DeleteVehicle(1)
		_, err = vehicleHandler.VehicleById(1)
		require.NotNil(t, err)
		vehicles, err = vehicleHandler.Vehicles(1)
		require.NoError(t, err)
		require.Equal(t, 1, len(vehicles))
	})
}
