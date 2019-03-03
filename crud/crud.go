package crud

import (
	"context"
	"log"
	"net/http"

	"github.com/mongodb/mongo-go-driver/mongo/options"

	"github.com/AMCCG/project-2-backend-golang/structure"
	"github.com/AMCCG/project-2-backend-golang/utils"
	"github.com/mongodb/mongo-go-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindAll(entity interface{}) structure.ResponseApi {
	var response structure.ResponseApi
	mongoDB := Connect()
	var results []*structure.User

	collections := mongoDB.Database("golang").Collection(utils.Reflection(entity))
	// cur, err := collections.Aggregate(context.TODO(), entity, options.Aggregate())
	cur, err := collections.Find(context.TODO(), entity, options.Find())

	if err != nil {
		log.Print("query error ", err.Error())
		response.Status = "error"
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
	} else {
		for cur.Next(context.TODO()) {
			var user structure.User
			err := cur.Decode(&user)
			if err != nil {
				log.Fatal("Decode ", err)
			}
			results = append(results, &user)
		}
		response.Status = "ok"
		response.Code = http.StatusOK
		response.Content = results
	}
	defer CloseDatabase(mongoDB)
	return response
}

func Insert(entity interface{}) structure.ResponseApi {
	var response structure.ResponseApi
	mongoDB := Connect()

	collections := mongoDB.Database("golang").Collection(utils.Reflection(entity))
	insertResult, err := collections.InsertOne(context.TODO(), entity)
	if err != nil {
		log.Print(err.Error())
		response.Status = "error"
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
	} else {
		log.Print("Inserted a single document: ", insertResult.InsertedID)
		response.Status = "create success"
		response.Code = http.StatusCreated
	}
	defer CloseDatabase(mongoDB)
	return response
}

func Update(entity interface{}) structure.ResponseApi {
	var response structure.ResponseApi
	mongoDB := Connect()

	filter := bson.M{"firstname": "Apisit"}
	update := bson.M{"$set": bson.M{"firstname": "Tee"}}

	collections := mongoDB.Database("golang").Collection(utils.Reflection(entity))
	updateResult, err := collections.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Print(err.Error())
		response.Status = "error"
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
	} else {
		log.Print("MatchedCount : ", updateResult.MatchedCount)
		log.Print("ModifiedCount : ", updateResult.ModifiedCount)
		log.Print("UpsertedID : ", updateResult.UpsertedID)
		response.Status = "create success"
		response.Code = http.StatusOK
	}
	defer CloseDatabase(mongoDB)
	return response
}

func Connect() *mongo.Client {
	mongoDB, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")
	if err != nil {
		log.Fatal("can't connect database ", err)
	}
	log.Print("Connect Success")
	return mongoDB
}

func CloseDatabase(database *mongo.Client) {
	err := database.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Connection to MongoDB closed.")
}