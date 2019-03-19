package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"project-2-backend-golang/structure"
	"strings"
)

func JWTEncoded(entity structure.User) string {
	var token string
	if entity.FIRSTNAME != "" {
		token = generateToken(entity)
	}
	return token
}

type HEADER struct {
	Agl string `json:"alg"`
	Typ string `json:"typ"`
}

func generateToken(user structure.User) string {
	header := HEADER{Agl: "HS256", Typ: "JWT"}
	var u structure.User
	u.FIRSTNAME = "apisit"
	jsonHeader, err := json.Marshal(header)
	jsonPayLoad, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error ", err.Error())
	}
	fmt.Println("json header : ", string(jsonHeader))
	fmt.Println("json user : ", string(jsonPayLoad))
	fmt.Println("json user remove field : ", removeStringJson(string(jsonPayLoad)))
	var jwt structure.JWT
	jwt.HEADER = base64.StdEncoding.EncodeToString([]byte(jsonHeader))
	jwt.PAYLOAD = base64.StdEncoding.EncodeToString([]byte(jsonPayLoad))
	fmt.Println("Header : ", jwt.HEADER)
	fmt.Println("PAYLOAD : ", jwt.PAYLOAD)
	// jwt.SIGNATURE = ""
	return jwt.HEADER + "." + jwt.PAYLOAD
}

func removeStringJson(json string) string {
	var jsonAfter string = "{"
	var field []string
	field = strings.Split(json, ",")
	fmt.Println(field)
	for i, t := range field {
		replace := strings.NewReplacer("{", "", "}", "")
		t = replace.Replace(t)
		fmt.Println(i, t)
		if strings.Index(t, "\"\"") == -1 && strings.Index(t, "\"\"") == -1 {
			jsonAfter = jsonAfter + t + ","
		}
	}
	fmt.Println("jsonAfter[0:] ", jsonAfter[0:len(jsonAfter)-1]+"}")
	return jsonAfter[0:len(jsonAfter)-1] + "}"
}
