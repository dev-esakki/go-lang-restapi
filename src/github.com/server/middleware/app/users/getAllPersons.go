package users


import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	Person "github.com/server/modals/person"
)

func GetAllUsers() ([]Person.Person) {
	data, err := collection.Find(context.TODO(), bson.D{{}}) //or bson.M
	if err != nil {
		log.Fatal(err)
	}
	var persons []Person.Person
	if err = data.All(context.TODO(), &persons); err != nil {
		log.Fatal(err)
	}
	return persons
}