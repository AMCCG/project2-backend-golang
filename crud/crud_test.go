package crud

import (
	"net/http"
	"testing"

)

func TestInsert(t *testing.T) {
	user := structure.User{FirstName: "Apisti", LastName: "Amornchanchaigul"}
	response := Insert(user)
	if response.Code != http.StatusCreated {
		t.Error("Insert Failed")
	}
}

func TestFindAll(t *testing.T) {
	user := structure.User{FirstName: "Apisti", LastName: "Amornchanchaigul"}
	response := FindAll(user)
	if response.Code != http.StatusOK {
		t.Error("Find Failed ", response.Code)
	}
	if response.Content == nil {
		t.Error("Content null ", response.Content)
	}
}

func TestUpdate(t *testing.T) {
	user := structure.User{FirstName: "Apisti", LastName: "Amornchanchaigul"}
	response := Update(user)
	if response.Code != http.StatusOK {
		t.Errorf("Update Failed %d message %s ", response.Code, response.Message)
	}
}
