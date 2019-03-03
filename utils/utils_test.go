package utils

import (
	"testing"

	"github.com/AMCCG/project-2-backend-golang/structure"
)

func TestReflection(t *testing.T) {
	var user structure.User
	if result := Reflection(user); result != "user" {
		t.Error("reflection failed ", result)
	}
}
