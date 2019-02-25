package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AMCCG/project-2-backend-golang/constant"
	"github.com/AMCCG/project-2-backend-golang/structure"
)

func TestGetAllUsers(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, constant.UsersURL, nil)
	if err != nil {
		t.Fatal(err)
	}
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(UserHandler)
	handler.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Errorf("Status code for /users is wrong. Have: %d, want: %d.", response.Code, http.StatusOK)
	}
	var responseApi structure.ResponseApi
	error := json.Unmarshal([]byte(response.Body.String()), &responseApi)
	if error != nil {
		t.Fatal(error)
	}

	x := responseApi.Content
	t.Error("X :  ", x)

	if responseApi.Code != http.StatusOK {
		t.Error("ResponseApi :  ", responseApi)
		t.Errorf("handler returned unexpected body: got %v want %v ",
			response.Body.String(), http.StatusOK)
	}
}

func TestFindById(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, constant.UsersURL+"1", nil)
	if err != nil {
		t.Fatal(err)
	}
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(UserHandler)
	handler.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Errorf("Status code for /users is wrong. Have: %d, want: %d.", response.Code, http.StatusOK)
	}
	var responseApi structure.ResponseApi
	error := json.Unmarshal([]byte(response.Body.String()), &responseApi)
	if error != nil {
		t.Fatal(error)
	}
	if responseApi.Code != http.StatusOK {
		t.Error("ResponseApi :  ", responseApi)
		t.Errorf("handler returned unexpected body: got %v want %v ",
			response.Body.String(), http.StatusCreated)
	}
}

func TestAddUsers(t *testing.T) {
	request, err := http.NewRequest(http.MethodPost, constant.UsersURL, nil)
	if err != nil {
		t.Fatal(err)
	}
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(UserHandler)
	handler.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Errorf("Status code for /users is wrong. Have: %d, want: %d.", response.Code, http.StatusOK)
	}
	var responseApi structure.ResponseApi
	error := json.Unmarshal([]byte(response.Body.String()), &responseApi)
	if error != nil {
		t.Fatal(error)
	}
	if responseApi.Code != http.StatusCreated {
		t.Error("ResponseApi :  ", responseApi)
		t.Errorf("handler returned unexpected body: got %v want %v ",
			response.Body.String(), http.StatusCreated)
	}
}
func TestUpdateUsers(t *testing.T) {
	request, err := http.NewRequest(http.MethodPut, constant.UsersURL, nil)
	if err != nil {
		t.Fatal(err)
	}
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(UserHandler)
	handler.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Errorf("Status code for /users is wrong. Have: %d, want: %d.", response.Code, http.StatusOK)
	}
	var responseApi structure.ResponseApi
	error := json.Unmarshal([]byte(response.Body.String()), &responseApi)
	if error != nil {
		t.Fatal(error)
	}
	if responseApi.Code != http.StatusOK {
		t.Error("ResponseApi :  ", responseApi)
		t.Errorf("handler returned unexpected body: got %v want %v ",
			response.Body.String(), http.StatusCreated)
	}
}

func TestDeleteUsers(t *testing.T) {
	request, err := http.NewRequest(http.MethodDelete, constant.UsersURL, nil)
	if err != nil {
		t.Fatal(err)
	}
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(UserHandler)
	handler.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Errorf("Status code for /users is wrong. Have: %d, want: %d.", response.Code, http.StatusOK)
	}
	var responseApi structure.ResponseApi
	error := json.Unmarshal([]byte(response.Body.String()), &responseApi)
	if error != nil {
		t.Fatal(error)
	}
	if responseApi.Code != http.StatusOK {
		t.Error("ResponseApi :  ", responseApi)
		t.Errorf("handler returned unexpected body: got %v want %v ",
			response.Body.String(), http.StatusCreated)
	}
}
