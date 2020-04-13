package router


import (
	"github.com/gorilla/mux"
	"github.com/server/middleware/middleperson"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	r:= mux.NewRouter()

	r.HandleFunc("/", middleperson.CreatePerson).Methods("GET", "OPTIONS")
	r.HandleFunc("/getPerson", middleperson.GetPerson).Methods("GET", "OPTIONS")
	r.HandleFunc("/getAllPerson", middleperson.GetAllPersons).Methods("GET", "OPTIONS")
	r.HandleFunc("/CreateManyPerson", middleperson.CreateManyPerson).Methods("GET", "OPTIONS")
	r.HandleFunc("/UpdatePerson", middleperson.UpdatePerson).Methods("GET", "OPTIONS")
	
	return r
}