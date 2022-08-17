package service

import (
	"CRUD/database"
	"CRUD/model"
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

var validate = validator.New()

func GetAllContacts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var contacts []model.Contact
	err := database.GetAllContacs(&contacts)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("500 - Error...")
		return
	}
	json.NewEncoder(response).Encode(contacts)
}

//Get Specific Contact detail
func GetContact(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	var contact model.Contact
	//Validation
	err := validate.Struct(contact)
	if err != nil {
		fmt.Println("Invalid input", err)
	}
	err = validate.Struct(params)
	if err != nil {
		fmt.Println("Invalid input", err)
	}
	err = database.GetContact(&contact, params)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("500 - Error...")
		return
	}
	json.NewEncoder(response).Encode(contact)

}

//Creating Contacts
func CreateContact(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/")
	var contact model.Contact

	json.NewDecoder(request.Body).Decode(&contact)

	//Validation
	err := validate.Struct(contact)
	if err != nil {
		fmt.Println("Invalid input", err)
	}
	err = database.CreateContact(&contact)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("500 - Error...")
		return

	}
	json.NewEncoder(response).Encode(contact)
}

// Update/Modify
func UpdateContact(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	var contact model.Contact

	//Validation
	err := validate.Struct(contact)
	if err != nil {
		fmt.Println("Invalid input", err)
	}
	err = validate.Struct(params)
	if err != nil {
		fmt.Println("Invalid input", err)
	}
	err = database.UpdateContact(&contact, params)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("500 - Error...")
		return
	}
	json.NewEncoder(response).Encode(contact)
}

//Deleting
func DeleteContact(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	err := validate.Struct(params)
	if err != nil {
		fmt.Println("Invalid input", err)
	}
	var contact model.Contact
	err = database.DeleteContact(&contact, params)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("500 - Error...")
		return
	}
	json.NewEncoder(response).Encode("The User is Deleted ...")
}
