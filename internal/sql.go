package internal

import (
	"awesomeProject19/entities"
	"awesomeProject19/logger"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func SavePerson(db *gorm.DB, people *entities.Person) error {
	logger.Log.Infof("Attempting to save person: %+v\n", people)
	if err := db.Create(people).Error; err != nil {
		logger.Log.Errorf("Error saving person: %v\n", err)
		return err
	}
	logger.Log.Info("Person saved successfully.")
	return nil
}

func DeletePerson(db *gorm.DB, id uint) error {
	logger.Log.Infof("Attempting to delete person with ID: %d\n", id)
	var person entities.Person
	res := db.Delete(&person, id)

	if res.Error != nil {
		logger.Log.Errorf("Error deleting person: %v\n", res.Error)
		return res.Error
	}

	if res.RowsAffected == 0 {
		err := fmt.Errorf("person with id %d not found", id)
		logger.Log.Warn(err.Error())
		return err
	}

	logger.Log.Info("Person deleted successfully.")
	return nil
}

func UpdatePerson(db *gorm.DB, person *entities.Person) error {
	logger.Log.Infof("Attempting to update person with ID: %d\n", person.ID)
	existingPerson := &entities.Person{}
	result := db.First(existingPerson, person.ID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			err := fmt.Errorf("person with id %d not found", person.ID)
			logger.Log.Warn(err.Error())
			return err
		}
		logger.Log.Errorf("Error finding person: %v\n", result.Error)
		return result.Error
	}

	result = db.Model(existingPerson).Updates(person)

	if result.Error != nil {
		logger.Log.Errorf("Error updating person: %v\n", result.Error)
		return result.Error
	}

	logger.Log.Info("Person updated successfully.")
	return nil
}

func GetPerson(db *gorm.DB, name string, limit, offset int) ([]entities.Person, error) {
	logger.Log.Infof("Fetching persons with name containing: %s, limit: %d, offset: %d\n", name, limit, offset)
	var person []entities.Person

	result := db.Where("name ILIKE ?", "%"+name+"%").
		Limit(limit).
		Offset(offset).
		Find(&person)

	if result.Error != nil {
		logger.Log.Errorf("Error fetching persons: %v\n", result.Error)
		return nil, result.Error
	}

	logger.Log.Infof("Fetched %d persons successfully.\n", len(person))
	return person, nil
}

func GetPersonBySurname(db *gorm.DB, surname string, limit, offset int) ([]entities.Person, error) {
	logger.Log.Infof("Fetching persons with surname containing: %s, limit: %d, offset: %d\n", surname, limit, offset)
	var users []entities.Person

	query := db.Model(&entities.Person{})

	if surname != "" {
		query = query.Where("surname ILIKE ?", "%"+surname+"%")
	}

	query = query.Order("surname ASC").Limit(limit).Offset(offset)

	if err := query.Find(&users).Error; err != nil {
		logger.Log.Errorf("Error fetching persons by surname: %v\n", err)
		return nil, err
	}

	logger.Log.Infof("Fetched %d persons by surname successfully.\n", len(users))
	return users, nil
}

func GetPersonByName(db *gorm.DB, name string, limit, offset int) ([]entities.Person, error) {
	logger.Log.Infof("Fetching persons with name containing: %s, limit: %d, offset: %d\n", name, limit, offset)
	var users []entities.Person

	query := db.Model(&entities.Person{})

	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	query = query.Order("name ASC").Limit(limit).Offset(offset)

	if err := query.Find(&users).Error; err != nil {
		logger.Log.Errorf("Error fetching persons by name: %v\n", err)
		return nil, err
	}

	logger.Log.Infof("Fetched %d persons by name successfully.\n", len(users))
	return users, nil
}
