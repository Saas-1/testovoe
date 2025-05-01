package main

import (
	"awesomeProject19/internal"
	"log"
	"net/http"
)

// @title Example API
// @version 1.0
// @description This is a simple API for managing people.
// @termsOfService http://example.com/terms/
// @contact.name Support
// @contact.url http://example.com/support
// @contact.email support@example.com
// @license.name MIT
// @license.url http://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /
func main() {
	_, err := internal.ConnectDB()
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	http.HandleFunc("/add-person", internal.AllowMethod(internal.AddHuman, http.MethodPost))
	http.HandleFunc("/delete-person", internal.AllowMethod(internal.DeleteHuman, http.MethodDelete))
	http.HandleFunc("/update-person", internal.AllowMethod(internal.UpdateHuman, http.MethodPut))
	http.HandleFunc("/name", internal.AllowMethod(internal.GetHumanByName, http.MethodGet))
	http.HandleFunc("/ascname", internal.AllowMethod(internal.GetAscFromName, http.MethodGet))
	http.HandleFunc("/ascsurname", internal.AllowMethod(internal.GetAscFromSurname, http.MethodGet))

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
