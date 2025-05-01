package internal

import (
	"awesomeProject19/entities"
	"awesomeProject19/logger"
	"encoding/json"
	"net/http"
	"strconv"
)

// AddHuman godoc
// @Summary Add a person
// @Description Add a new person to the database
// @Accept  json
// @Produce  json
// @Param person body entities.FIO true "Person data"
// @Success 200 {object} entities.Person
// @Failure 400 {string} string "Error decoding JSON"
// @Failure 500 {string} string "Error saving person"
// @Router /add-person [post]
func AddHuman(w http.ResponseWriter, r *http.Request) {
	var fio entities.FIO
	err := json.NewDecoder(r.Body).Decode(&fio)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	logger.Log.Infof("Adding person: %+v", fio)

	age, _ := addAge(fio.Name)
	gender, _ := addGender(fio.Name)
	nationality, _ := addNationality(fio.Name)

	people := entities.Person{
		Name:        fio.Name,
		Surname:     fio.Surname,
		Patronymic:  fio.Patronymic,
		Gender:      gender,
		Age:         age,
		Nationality: nationality,
	}

	err = SavePerson(db, &people)
	if err != nil {
		http.Error(w, "Error saving person", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(people); err != nil {
		logger.Log.Errorf("Error encoding JSON: %v", err)
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

// DeleteHuman godoc
// @Summary Delete a person
// @Description Delete a person from the database by ID
// @Param id query int true "Person ID"
// @Success 200 {string} string "Person deleted successfully"
// @Failure 400 {string} string "ID not found in query parameters"
// @Failure 500 {string} string "Failed to delete person"
// @Router /delete-person [delete]
func DeleteHuman(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID not found in query parameters", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	logger.Log.Infof("Deleting person with ID: %d", id)
	err = DeletePerson(db, uint(id))
	if err != nil {
		http.Error(w, "Failed to delete person", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	logger.Log.Infof("Person with ID: %d deleted successfully", id)
}

// UpdateHuman godoc
// @Summary Update a person
// @Description Update a person's details in the database
// @Param id query int true "Person ID"
// @Param person body entities.Person true "Updated person data"
// @Success 200 {string} string "Person updated successfully"
// @Failure 400 {string} string "Invalid request body"
// @Failure 500 {string} string "Failed to update person"
// @Router /update-person [put]
func UpdateHuman(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID not found in query parameters", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var person entities.Person
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	person.ID = uint(id)
	if err := UpdatePerson(db, &person); err != nil {
		http.Error(w, "Failed to update person", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	logger.Log.Infof("Person with ID: %d updated successfully", id)
}

// GetHumanByName godoc
// @Summary Get people by name
// @Description Retrieve people from the database by name
// @Param name query string true "Name of the person"
// @Success 200 {array} entities.Person
// @Failure 500 {string} string "Database error"
// @Router /name [get]
func GetHumanByName(w http.ResponseWriter, r *http.Request) {
	limit := 1
	offset := 0
	name := r.URL.Query().Get("name")
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	if limitStr != "" {
		l, err := strconv.Atoi(limitStr)
		if err == nil {
			limit = l
		}
	}
	if offsetStr != "" {
		o, err := strconv.Atoi(offsetStr)
		if err == nil {
			offset = o
		}
	}

	logger.Log.Infof("Retrieving people with name: %s, limit: %d, offset: %d", name, limit, offset)
	people, err := GetPerson(db, name, limit, offset)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(people); err != nil {
		logger.Log.Errorf("Error encoding JSON: %v", err)
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

// GetAscFromName godoc
// @Summary Get people by surname
// @Description Retrieve people by surname
// @Param surname query string true "Surname"
// @Success 200 {array} entities.Person
// @Failure 500 {string} string "Database error"
// @Router /ascname [get]
func GetAscFromName(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("surname")

	limit := 1
	offset := 0

	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = l
		}
	}
	if offsetStr := r.URL.Query().Get("offset"); offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil {
			offset = o
		}
	}

	logger.Log.Infof("Retrieving people by surname: %s, limit: %d, offset: %d", firstName, limit, offset)
	people, err := GetPersonByName(db, firstName, limit, offset)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(people); err != nil {
		logger.Log.Errorf("Error encoding JSON: %v", err)
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

// GetAscFromSurname godoc
// @Summary Get people by first name
// @Description Retrieve people by first name
// @Param name query string true "First name"
// @Success 200 {array} entities.Person
// @Failure 500 {string} string "Database error"
// @Router /ascsurname [get]
func GetAscFromSurname(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("name")

	limit := 1
	offset := 0

	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = l
		}
	}
	if offsetStr := r.URL.Query().Get("offset"); offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil {
			offset = o
		}
	}

	logger.Log.Infof("Retrieving people by name: %s, limit: %d, offset: %d", firstName, limit, offset)
	people, err := GetPersonBySurname(db, firstName, limit, offset)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(people); err != nil {
		logger.Log.Errorf("Error encoding JSON: %v", err)
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}
