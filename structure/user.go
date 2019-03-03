package structure

type User struct {
	// Id        bson.ObjectId `bson:"_id,omitempty"`
	Id        string `bson:"_id,omitempty"`
	FirstName string
	LastName  string
}
