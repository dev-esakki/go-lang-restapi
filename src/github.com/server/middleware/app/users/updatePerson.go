package users

import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SUserUpdate struct {
	id int  `json:"id"`
    name string `json:"name"`
    age  int `json:"age"`
    city string `json:"city"`
}

func (input SUserUpdate) UserUpdate() (mongo.UpdateResult) {	
	filter := bson.D{{"id", input.id}}
	newName := bson.D{
		{"$set", bson.D{
			{"name", input.name},
			{"age", input.age},
			{"city", input.city},
		}},
	}
	data, err := collection.UpdateOne(context.TODO(), filter, newName) 
	if err != nil {
		log.Fatal(err)
	}

	updatedObject := *data
	return updatedObject
}