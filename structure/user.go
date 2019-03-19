package structure

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FIRSTNAME string             `json:"firstname"`
	LASTNAME  string             `json:"lastname"`
	EMAIL     string             `json:"email"`
	PASSWORD  string             `json:"password"`
}
