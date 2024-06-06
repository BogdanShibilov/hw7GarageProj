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
		_ = user.Add(0, "John")
		_ = user.Add(1, "Jack")
		err := user.Add(2, "Janny")

		require.NoError(t, err)
		require.Equal(t, 3, goMap.GetGarage().UserCount())

		user.Delete(0)
		require.Equal(t, 2, goMap.GetGarage().UserCount())

		user.Delete(5)
		require.Equal(t, 2, goMap.GetGarage().UserCount())

		err = user.Update(2, "Johny")
		require.NoError(t, err)
		johnyUser, _ := user.ById(2)
		require.NotNil(t, johnyUser)
		require.Equal(t, "Johny", johnyUser.Name)

		err = user.Add(2, "Jennyfer")
		require.NotNil(t, err)
		require.Equal(t, 2, goMap.GetGarage().UserCount())

		err = user.Update(999, "Jest")
		require.NotNil(t, err)

		err = vehicle.Add(0, "vName1", "vBrand1", "vModel1", 1)
		require.NoError(t, err)

		vehicles, err := vehicle.AllBelongingTo(1)
		require.NoError(t, err)
		require.Equal(t, 1, len(vehicles))

		err = vehicle.Add(0, "vName1", "vBrand1", "vModel1", 1)
		require.NotNil(t, err)

		vehicles, err = vehicle.AllBelongingTo(1)
		require.NoError(t, err)
		require.Equal(t, 1, len(vehicles))

		err = vehicle.Add(1, "vName2", "vBrand2", "vModel2", 1)
		require.NoError(t, err)
		vehicles, err = vehicle.AllBelongingTo(1)
		require.NoError(t, err)
		require.Equal(t, 2, len(vehicles))

		v1, err := vehicle.ById(1)
		require.NoError(t, err)
		require.Equal(t, "vName2", v1.Name)

		err = vehicle.Update(1, "updatedVName2", "vBrand2", "vModel2")
		require.NoError(t, err)

		updatedV1, err := vehicle.ById(1)
		require.NoError(t, err)
		require.Equal(t, "updatedVName2", updatedV1.Name)

		err = vehicle.Update(999, "no", "no", "no")
		require.NotNil(t, err)

		vehicle.Delete(1)
		_, err = vehicle.ById(1)
		require.NotNil(t, err)
		vehicles, err = vehicle.AllBelongingTo(1)
		require.NoError(t, err)
		require.Equal(t, 1, len(vehicles))
	})
}
