package middleperson


import (
	"context"
	"log"
	"fmt"
	"github.com/server/db"
	Person "github.com/server/modals/person"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)
var client *mongo.Database = db.Database()

// Insert one task in the DB
//task 	Person.Person
func CreatePerson(w http.ResponseWriter, r *http.Request) {

	collection := client.Collection("persons")
	
	ruan := Person.Person{"qwert", 28, "urkad"}

	insertResult, err := collection.InsertOne(context.TODO(), ruan)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Document: ", insertResult.InsertedID) 
}