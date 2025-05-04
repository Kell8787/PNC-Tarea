package services

import (
	"errors"
	"github.com/google/uuid"
	"layersapi/entities"
	"layersapi/entities/dto"
	"layersapi/repositories"
	"regexp"
	"time"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (u UserService) Update(id string, user dto.UpdateUser) error {
	if len(user.Name) == 0 {
		return errors.New("name cannot be empty")
	}

	re, _ := regexp.Compile(`^[A-Za-z]+$`)
	if !re.MatchString(user.Name) {
		return errors.New("name must only contain alphabetic characters")

	}

	re, _ = regexp.Compile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

	if !re.MatchString(user.Email) {
		return errors.New("invalid email address")
	}

	err := u.userRepository.Update(id, user.Name, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (u UserService) Delete(id string) error {
	if len(id) == 0 {
		return errors.New("id cannot be empty")
	}

	_, err := u.userRepository.GetById(id)
	if err != nil {
		return errors.New("user not found")
	}

	err = u.userRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (u UserService) Create(user dto.CreateUser) error {
	if len(user.Name) == 0 {
		return errors.New("name cannot be empty")
	}

	re, _ := regexp.Compile(`^[A-Za-z]+$`)
	if !re.MatchString(user.Name) {
		return errors.New("name must only contain alphabetic characters")

	}

	re, _ = regexp.Compile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

	if !re.MatchString(user.Email) {
		return errors.New("invalid email address")
	}

	id, _ := uuid.NewUUID()
	meta := entities.Metadata{
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
		CreatedBy: "webapp",
	}
	newUser := entities.NewUser(id.String(), user.Name, user.Email, meta)

	err := u.userRepository.Create(newUser)
	if err != nil {
		return err
	}
	return nil
}

func (u UserService) GetAll() ([]entities.User, error) {
	data, err := u.userRepository.GetAll()
	if err != nil {
		return []entities.User{}, err
	}

	return data, nil
}

func (u UserService) GetById(id string) (entities.User, error) {
	if len(id) == 0 {
		return entities.User{}, errors.New("id cannot be empty")
	}

	user, err := u.userRepository.GetById(id)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}
