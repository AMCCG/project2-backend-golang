package utils

import (
	"project-2-backend-golang/structure"
	"testing"
)

func TestEncoded(t *testing.T) {
	var user structure.User
	user.FIRSTNAME = "apisit"
	// var token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJGSVJTVE5BTUUiOiJhcGlzaXQifQ.MbMJUL9r4Qq6CKOrE5oI4B9Xr-Gbxn17tajWhbiPQz8"
	if result := JWTEncoded(user); result != "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"+"."+"eyJmaXJzdG5hbWUiOiJhcGlzaXQifQ" {
		t.Error("Token invalid")
	}
}
