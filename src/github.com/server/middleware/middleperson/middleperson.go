package middleperson


import (
	"context"
	"log"
	"fmt"
	"github.com/server/db"
	Person "github.com/server/modals/person"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"encoding/json"

)
var client *mongo.Database = db.Database()
var collection *mongo.Collection = client.Collection("persons")

// Insert one task in the DB
//task 	Person.Person
func CreatePerson(w http.ResponseWriter, r *http.Request) {

	
	ruan := Person.Person{"qwert", 28, "urkad"}

	insertResult, err := collection.InsertOne(context.TODO(), ruan)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Document: ", insertResult.InsertedID) 
}


func CreateManyPerson(w http.ResponseWriter, r *http.Request) {

    ruan := Person.Person{"Ruan", 34, "Cape Town"}
    james := Person.Person{"James", 32, "Nairobi"}
    frankie := Person.Person{"Frankie", 31, "Nairobi"}

    trainers := []interface{}{james, frankie, ruan}

    insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
    if err != nil {
      log.Fatal(err)
  }
    fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
}

func GetAllPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	data, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var persons []bson.M
	if err = data.All(context.TODO(), &persons); err != nil {
		log.Fatal(err)
	}
	fmt.Println(persons)
	json.NewEncoder(w).Encode(persons)
 
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	filter := bson.D{}
	var result Person.Person

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

  fmt.Printf("Found a single document: %+v\n", result)

  findOptions := options.Find()
    findOptions.SetLimit(2)
}