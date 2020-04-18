package middleperson


import (
	"context"
	"log"
	"fmt"
	Person "github.com/server/modals/person"
	StructPlaces "github.com/server/modals/places"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"encoding/json"

)

var collection *mongo.Collection = Person.PersonsDb()

func CreatePerson(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	  
    var t Person.Person
	_ = json.NewDecoder(r.Body).Decode(&t)
	fmt.Println(t)

	insertResult, err := collection.InsertOne(context.TODO(), t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Document: ", insertResult.InsertedID) 
	json.NewEncoder(w).Encode(t)
}


func CreateManyPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    /* ruan := Person.Person{"Ruan", 34, "Cape Town"}
    james := Person.Person{"James", 32, "Nairobi"}
	frankie := Person.Person{"Frankie", 31, "Nairobi"} */
	var t []Person.Person
	_ = json.NewDecoder(r.Body).Decode(&t)
	fmt.Println(t)
	var ui []interface{}
	//trainers := []interface{}{ruan, james}
	for _, t := range t {
		ui = append(ui, t)
	}

    insertManyResult, err := collection.InsertMany(context.TODO(), ui)
    if err != nil {
      log.Fatal(err)
    }
	fmt.Println("Inserted multiple documents: ", insertManyResult)
	json.NewEncoder(w).Encode(ui)
}

func GetAllPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	data, err := collection.Find(context.TODO(), bson.D{{}}) //or bson.M
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
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	filter := bson.D{{"name", "Ruan"}}
	//filter := bson.M{{"name": "Ruan"}}
	var result Person.Person

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

  fmt.Printf("Found a single document: %+v\n", result)
  json.NewEncoder(w).Encode(result)

  findOptions := options.Find()
	findOptions.SetLimit(2)	
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	filter := bson.D{{"name", "Ruan"}}
	newName := bson.D{
		{"$set", bson.D{
			{"name", "Updated Name of person 1"},
		}},
	}
	data, err := collection.UpdateOne(context.TODO(), filter, newName) 
	if err != nil {
		log.Fatal(err)
	}

	updatedObject := *data
	json.NewEncoder(w).Encode(updatedObject)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	filter := bson.D{{"name", "Updated Name of person 1"}}
	data, err := collection.DeleteOne(context.TODO(), filter) 
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(data)
}

type UserPlaces struct {

	Age int `bson:"age"`
	City string `bson:"city"`
	Id int `bson:"id"`
	Places []StructPlaces.Places `bson:"places"`
}

func GetUserPlaces(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	matchStage := bson.D{{"$match", bson.D{{"id", 1}}}}
	lookupStage := bson.D{{"$lookup", bson.D{{"from", "places"}, {"localField", "id"}, {"foreignField", "userid"}, {"as", "Places"}}}}
	//unwindStage := bson.D{{"$unwind", bson.D{{"path", "$podcast"}, {"preserveNullAndEmptyArrays", false}}}}

	showLoadedCursor, err := collection.Aggregate(context.TODO(), mongo.Pipeline{matchStage, lookupStage})
	if err != nil {
		panic(err)
	}
	var showsLoaded []UserPlaces  //use struct or bson.M
	if err = showLoadedCursor.All(context.TODO(), &showsLoaded); err != nil {
		panic(err)
	}
	fmt.Println(showsLoaded)
	json.NewEncoder(w).Encode(showsLoaded)
}