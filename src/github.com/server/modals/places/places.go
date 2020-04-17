package places

import (
    "github.com/server/db"
    "go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Database = db.Database()
var collection *mongo.Collection

type Places struct {
    Id int `json: "id"`
    Userid int `json: "userid"`
    Note string `json: "note"`
    City string `json: "city"`
}

func PlacesDb()*mongo.Collection {
    collection  = client.Collection("places")
    return collection
}
