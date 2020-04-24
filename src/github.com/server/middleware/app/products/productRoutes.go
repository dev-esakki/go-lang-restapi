package products


import (
	"context"
	"log"
	"fmt"
	Places "github.com/server/modals/places"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"encoding/json"

)

var colPlaces *mongo.Collection = Places.PlacesDb()


func CreateManyPlaces(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var t []Places.Places
	_ = json.NewDecoder(r.Body).Decode(&t)
	fmt.Println(t)
	var ui []interface{}	
	for _, t := range t {
		ui = append(ui, t)
	}

    insertManyResult, err := colPlaces.InsertMany(context.TODO(), ui)
    if err != nil {
      log.Fatal(err)
    }
	fmt.Println("Inserted multiple places: ", insertManyResult)
	json.NewEncoder(w).Encode(ui)
}

func GetAllPlaces( w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	data, err := colPlaces.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var places []bson.M
	if err = data.All(context.TODO(), &places); err != nil {
		log.Fatal(err)
	}
	fmt.Println(places)
	json.NewEncoder(w).Encode(places)
}

