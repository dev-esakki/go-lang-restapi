package person

import (
    "github.com/server/db"
    "go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Database = db.Database()
var collection *mongo.Collection

type Person struct {
    Id int `json: "id"`
    Name string `json: "name"`
    Age  int `json: "age"`
    City string `json: "city"`
}

func PersonsDb()*mongo.Collection {
    collection  = client.Collection("persons")
    return collection
}
