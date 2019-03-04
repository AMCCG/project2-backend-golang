package structure

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FIRSTNAME string
	LASTNAME  string
	EMAIL     string
	PASSWORD  string
}
