package csv

import (
	"encoding/csv"
	"errors"
	"layersapi/entities"
	"os"
	"path/filepath"
	"time"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u UserRepository) GetAll() ([]entities.User, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []entities.User{}, err
	}

	filePath := filepath.Join(cwd, "data", "data.csv")
	file, err := os.Open(filePath)
	if err != nil {
		return []entities.User{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return []entities.User{}, err
	}

	var result []entities.User

	for i, record := range records {
		if i == 0 {
			continue
		}

		createdAt, _ := time.Parse(time.RFC3339, record[3])
		updatedAt, _ := time.Parse(time.RFC3339, record[4])
		meta := entities.Metadata{
			CreatedAt: createdAt.String(),
			UpdatedAt: updatedAt.String(),
			CreatedBy: record[5],
			UpdatedBy: record[6],
		}
		result = append(result, entities.NewUser(record[0], record[1], record[2], meta))
	}

	return result, nil
}

func (u UserRepository) GetById(id string) (entities.User, error) {
	cwd, err := os.Getwd()

	filePath := filepath.Join(cwd, "data", "data.csv")
	file, err := os.Open(filePath)
	if err != nil {
		return entities.User{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return entities.User{}, err
	}

	for i, record := range records {
		if i == 0 {
			continue
		} else if record[0] == id {

			createdAt, _ := time.Parse(time.RFC3339, record[3])
			updatedAt, _ := time.Parse(time.RFC3339, record[4])
			meta := entities.Metadata{
				CreatedAt: createdAt.String(),
				UpdatedAt: updatedAt.String(),
				CreatedBy: record[5],
				UpdatedBy: record[6],
			}
			return entities.NewUser(record[0], record[1], record[2], meta), nil
		}

	}

	return entities.User{}, errors.New("user not found")
}

func (u UserRepository) Create(user entities.User) error {
	cwd, err := os.Getwd()

	filePath := filepath.Join(cwd, "data", "data.csv")
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	newUser := []string{
		user.Id,
		user.Name,
		user.Email,
		user.Metadata.CreatedAt,
		user.Metadata.UpdatedAt,
		"webapp",
		"webapp",
	}

	if err := writer.Write(newUser); err != nil {
		return err
	}

	return nil
}

func (u UserRepository) Update(id, name, email string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	filePath := filepath.Join(cwd, "data", "data.csv")
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	file.Close()
	if err != nil {
		return err
	}

	found := false
	for i, record := range records {
		if i == 0 {
			continue
		}
		if record[0] == id {
			record[1] = name
			record[2] = email
			record[4] = time.Now().Format(time.RFC3339)
			record[6] = "webapp"
			records[i] = record
			found = true
			break
		}
	}
	if !found {
		return errors.New("user not found")
	}

	tempFilePath := filepath.Join(cwd, "data", "data_temp.csv")
	tempFile, err := os.Create(tempFilePath)
	if err != nil {
		return err
	}

	writer := csv.NewWriter(tempFile)
	if err := writer.WriteAll(records); err != nil {
		tempFile.Close()
		return err
	}
	writer.Flush()
	tempFile.Close()

	time.Sleep(50 * time.Millisecond)

	if err := os.Remove(filePath); err != nil {
		return err
	}
	if err := os.Rename(tempFilePath, filePath); err != nil {
		return err
	}

	return nil
}

func (u UserRepository) Delete(id string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	filePath := filepath.Join(cwd, "data", "data.csv")

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	found := false
	updatedRecords := [][]string{}
	for i, record := range records {
		if i == 0 || record[0] != id {
			updatedRecords = append(updatedRecords, record)
		} else {
			found = true
		}
	}

	if !found {
		return errors.New("user not found")
	}

	tempFilePath := filepath.Join(cwd, "data", "data_temp.csv")
	tempFile, err := os.Create(tempFilePath)
	if err != nil {
		return err
	}

	writer := csv.NewWriter(tempFile)
	defer func() {
		writer.Flush()
		tempFile.Close()
	}()

	if err := writer.WriteAll(updatedRecords); err != nil {
		return err
	}

	file.Close()
	tempFile.Close()

	if err := os.Rename(tempFilePath, filePath); err != nil {
		return err
	}

	return nil
}
