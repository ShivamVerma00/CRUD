package handelfunc

import (
	"CRUD/database"
	"CRUD/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

var validate = validator.New()

//Get All Users
func GetUsers(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var users []model.Contact
	err := database.GetUsers(&users)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("500 - Error...")
		return
	}
	json.NewEncoder(response).Encode(users)
}

//Get Specific Contact detail
func GetUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	var user model.Contact
	err := database.GetUser(&user, params)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("500 - Error...")
		return
	}
	json.NewEncoder(response).Encode(user)

}

//Creating Contacts
func CreateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/")
	var user model.Contact

	json.NewDecoder(request.Body).Decode(&user)

	//Validation
	eror := validate.Struct(user)
	if eror != nil {
		log.Panic("Invalid input", eror)
	}
	err := database.CreateUser(&user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("500 - Error...")
		return

	}
	json.NewEncoder(response).Encode(user)
}

// Update/Modify
func UpdateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	var user model.Contact

	//Validation
	eror := validate.Struct(user)
	if eror != nil {
		log.Panic("Invalid input", eror)
	}
	err := database.UpdateUser(&user, params)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("500 - Error...")
		return
	}
	json.NewEncoder(response).Encode(user)
}

//Deleting
func DeleteUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	var user model.Contact
	err := database.DeleteUser(&user, params)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("500 - Error...")
		return
	}
	json.NewEncoder(response).Encode("The User is Deleted ...")
}
