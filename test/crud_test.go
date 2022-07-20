package main

import (
	handler "CRUD/handlerFunc"
	"CRUD/model"
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-playground/validator/v10"
)

var reader = bufio.NewReader(os.Stdin)

var validate = validator.New()

func TestGetUsers(t *testing.T) {

	req, err := http.NewRequest("GET", "/get", nil)
	if err != nil {
		t.Fatal(err)
	}
	r := httptest.NewRecorder()
	control := http.HandlerFunc(handler.GetUsers)
	control.ServeHTTP(r, req)

	if status := r.Code; status != http.StatusOK {
		t.Errorf("Wrong answer %v  got %v", status, http.StatusOK)
	}

	expected := `[{"id":1,"firstname":"Shivam","lastname":"Verma","email":"ShivamVerma@gmail.com"},{"firstname":"TFT","lastname":"us","email":"TFTus@gmail.com"}]`
	if r.Body.String() != expected {
		t.Errorf("Wrong Answer By Handler: got %v want %v",
			r.Body.String(), expected)
	}
}

func TestGetUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/get", nil)
	if err != nil {
		t.Fatal(err)
	}
	r := httptest.NewRecorder()
	control := http.HandlerFunc(handler.GetUser)
	control.ServeHTTP(r, req)

	if status := r.Code; status != http.StatusOK {
		t.Errorf("Wrong answer %v  got %v", status, http.StatusOK)
	}

	expected := `{"id":1,"firstname":"Shivam","lastname":"Verma","email":"ShivamVerma@gmail.com"}`
	if r.Body.String() != expected {
		t.Errorf("Wrong Answer By Handler: got %v want %v",
			r.Body.String(), expected)
	}
}

func TestCreateUser(t *testing.T) {
	var jsonReq = []byte(`{"firstname":"Shivam","lastname":"Verma","email":"ShivamVerma@gmail.com"}`)

	req, err := http.NewRequest("POST", "/Create", bytes.NewBuffer(jsonReq))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	r := httptest.NewRecorder()

	control := http.HandlerFunc(handler.CreateUser)
	//dispatch the http request
	control.ServeHTTP(r, req)

	if status := r.Code; status != http.StatusOK {
		t.Errorf("Wrong answer %v  got %v", status, http.StatusOK)
	}

	//input
	var user model.Contact
	fmt.Println("Enter First Name: ")
	first, _ := reader.ReadString('\n')

	fmt.Println("Enter Last Name: ")
	last, _ := reader.ReadString('\n')

	fmt.Println("Enter Email: ")
	email, _ := reader.ReadString('\n')

	user.FirstName = first
	user.LastName = last
	user.EmailId = email

	eror := validate.Struct(user)
	if eror != nil {
		log.Panic("Invalid input", eror)
	}

	expected := `{"id":1,"firstname":"Shivam","lastname":"Verma","email":"ShivamVerma@gmail.com"}`
	if r.Body.String() != expected {
		t.Errorf("Wrong Answer By Handler: got %v want %v",
			r.Body.String(), expected)
	}
}

func TestUpdateUser(t *testing.T) {

	var jsonReq = []byte(`{"id":2,"firstname":"TFT","lastname":"USA","email":"TFTUSA@gmail.com"}`)

	req, err := http.NewRequest("PUT", "/update/{id}", bytes.NewBuffer(jsonReq))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	r := httptest.NewRecorder()
	control := http.HandlerFunc(handler.UpdateUser)
	control.ServeHTTP(r, req)
	if status := r.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//input
	var user model.Contact
	fmt.Println("Enter First Name: ")
	first, _ := reader.ReadString('\n')

	fmt.Println("Enter Last Name: ")
	last, _ := reader.ReadString('\n')

	fmt.Println("Enter Email: ")
	email, _ := reader.ReadString('\n')

	user.FirstName = first
	user.LastName = last
	user.EmailId = email

	eror := validate.Struct(user)
	if eror != nil {
		log.Panic("Invalid input", eror)
	}

	expected := `{"id":2,"firstname":"TFT","lastname":"USA","email":"TFTUSA@gmail.com"}`
	if r.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			r.Body.String(), expected)
	}
}

func TestDeleteUser(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/delete/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	r := httptest.NewRecorder()
	control := http.HandlerFunc(handler.DeleteUser)
	control.ServeHTTP(r, req)
	if status := r.Code; status != http.StatusOK {
		t.Errorf("Wrong answer %v  got %v", status, http.StatusOK)
	}

	expected := `{"id":1,"firstname":"Shivam","lastname":"Verma","email":"ShivamVerma@gmail.com"}`
	if r.Body.String() != expected {
		t.Errorf("Wrong Answer By Handler: got %v want %v",
			r.Body.String(), expected)
	}

}
