package users

import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	Person "github.com/server/modals/person"
)

type SGetUser struct {
	name string
}
 
func (input SGetUser) GetUser() (Person.Person) {
	filter := bson.D{{"name", input.name}}
	//filter := bson.M{{"name": "Ruan"}}
	var result Person.Person

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
    return result
}