package memory

import (
	"errors"
	"layersapi/data"
	"layersapi/entities"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u UserRepository) GetAll() ([]entities.User, error) {
	res := data.Data
	return res, nil
}

func (u UserRepository) GetById(id string) (entities.User, error) {
	for _, v := range data.Data {
		if v.Id == id {
			return v, nil
		}
	}

	return entities.User{}, errors.New("user not found")
}

func (u UserRepository) Create(user entities.User) error {
	data.Data = append(data.Data, user)
	return nil
}

func (u UserRepository) Update(id, name, email string) error {
	for i, v := range data.Data {
		if v.Id == id {
			data.Data[i].Name = name
			data.Data[i].Email = email
			return nil
		}
	}

	return errors.New("user not found")
}

func (u UserRepository) Delete(id string) error {
	for i, v := range data.Data {
		if v.Id == id {
			data.Data[i].Email = ""
			data.Data[i].Email = ""
		}
	}

	return errors.New("user not found")
}
