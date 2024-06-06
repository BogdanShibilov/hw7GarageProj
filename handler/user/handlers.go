package user

import (
	"hw7garageproj/model"
	"strings"
)

type Storage interface {
	AddUser(newUser *model.User)
	DeleteUser(id int)
	User(id int) *model.User
	UpdateUser(newUser *model.User)
}

func Add(id int, name string, to Storage) error {
	if id < 0 || strings.Trim(name, " ") == "" {
		return InvalidInputErr
	}

	if user := to.User(id); user != nil {
		return AlreadyExistsErr
	}

	newUser := model.User{
		Id:   id,
		Name: name,
	}

	to.AddUser(&newUser)
	return nil
}

func Delete(id int, from Storage) {
	if id < 0 {
		return
	}

	from.DeleteUser(id)
}

func ById(id int, from Storage) (*model.User, error) {
	if id < 0 {
		return nil, InvalidInputErr
	}

	if user := from.User(id); user != nil {
		return user, nil
	}
	return nil, NotFoundErr
}

func Update(userId int, newName string, in Storage) error {
	if userId < 0 || strings.Trim(newName, " ") == "" {
		return InvalidInputErr
	}

	var user *model.User
	if user = in.User(userId); user == nil {
		return NotFoundErr
	}
	user.Name = newName
	in.UpdateUser(user)
	return nil
}
