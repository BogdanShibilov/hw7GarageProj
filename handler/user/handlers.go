package user

import (
	"hw7garageproj/model"
	"hw7garageproj/storage/goMap"
	"strings"
)

var storage = goMap.GetGarage()

func Add(id int, name string) error {
	if id < 0 {
		return InvalidIdErr
	}
	if strings.Trim(name, " ") == "" {
		return InvalidNameErr
	}

	if user := storage.UserGet(id); user != nil {
		return AlreadyExistsErr
	}

	newUser := model.User{
		Id:   id,
		Name: name,
	}

	storage.AddUser(&newUser)
	return nil
}

func Delete(id int) {
	if id < 0 {
		return
	}

	storage.DeleteUser(id)
}

func ById(id int) (*model.User, error) {
	if id < 0 {
		return nil, InvalidIdErr
	}

	if user := storage.UserGet(id); user != nil {
		return user, nil
	}
	return nil, NotFoundErr
}

func Update(userId int, newName string) error {
	if userId < 0 {
		return InvalidIdErr
	}
	if strings.Trim(newName, " ") == "" {
		return InvalidNameErr
	}

	var user *model.User
	if user = storage.UserGet(userId); user == nil {
		return NotFoundErr
	}
	user.Name = newName
	storage.UpdateUser(user)
	return nil
}
