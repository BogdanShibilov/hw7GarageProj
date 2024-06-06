package test

import (
	"github.com/stretchr/testify/require"
	"hw7garageproj/handler/user"
	"hw7garageproj/handler/vehicle"
	"hw7garageproj/storage/goMap"
	"testing"
)

func TestList(t *testing.T) {
	t.Run("handlers", func(t *testing.T) {
		garage := goMap.NewGarage(10)

		_ = user.Add(0, "John", garage)
		_ = user.Add(1, "Jack", garage)
		err := user.Add(2, "Janny", garage)

		require.NoError(t, err)
		require.Equal(t, 3, garage.UserCount())

		user.Delete(0, garage)
		require.Equal(t, 2, garage.UserCount())

		user.Delete(5, garage)
		require.Equal(t, 2, garage.UserCount())

		err = user.Update(2, "Johny", garage)
		require.NoError(t, err)
		johnyUser, _ := user.ById(2, garage)
		require.NotNil(t, johnyUser)
		require.Equal(t, "Johny", johnyUser.Name)

		err = user.Add(2, "Jennyfer", garage)
		require.NotNil(t, err)
		require.Equal(t, 2, garage.UserCount())

		err = user.Update(999, "Jest", garage)
		require.NotNil(t, err)

		err = vehicle.Add(0, "vName1", "vBrand1", "vModel1", 1, garage, garage)
		require.NoError(t, err)

		vehicles, err := vehicle.AllBelongingTo(1, garage, garage)
		require.NoError(t, err)
		require.Equal(t, 1, len(vehicles))

		err = vehicle.Add(0, "vName1", "vBrand1", "vModel1", 1, garage, garage)
		require.NotNil(t, err)

		vehicles, err = vehicle.AllBelongingTo(1, garage, garage)
		require.NoError(t, err)
		require.Equal(t, 1, len(vehicles))

		err = vehicle.Add(1, "vName2", "vBrand2", "vModel2", 1, garage, garage)
		require.NoError(t, err)
		vehicles, err = vehicle.AllBelongingTo(1, garage, garage)
		require.NoError(t, err)
		require.Equal(t, 2, len(vehicles))

		v1, err := vehicle.ById(1, garage)
		require.NoError(t, err)
		require.Equal(t, "vName2", v1.Name)

		err = vehicle.Update(1, "updatedVName2", "vBrand2", "vModel2", garage)
		require.NoError(t, err)

		updatedV1, err := vehicle.ById(1, garage)
		require.NoError(t, err)
		require.Equal(t, "updatedVName2", updatedV1.Name)

		err = vehicle.Update(999, "no", "no", "no", garage)
		require.NotNil(t, err)

		vehicle.Delete(1, garage)
		_, err = vehicle.ById(1, garage)
		require.NotNil(t, err)
		vehicles, err = vehicle.AllBelongingTo(1, garage, garage)
		require.NoError(t, err)
		require.Equal(t, 1, len(vehicles))
	})
}
