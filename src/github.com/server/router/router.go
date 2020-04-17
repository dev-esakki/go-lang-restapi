package router


import (
	"github.com/gorilla/mux"
	middleware "github.com/server/middleware/middleperson"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	r:= mux.NewRouter()

	r.HandleFunc("/", middleware.CreatePerson).Methods("POST", "OPTIONS")
	r.HandleFunc("/getPerson", middleware.GetPerson).Methods("GET", "OPTIONS")
	r.HandleFunc("/getAllPerson", middleware.GetAllPersons).Methods("GET", "OPTIONS")
	r.HandleFunc("/CreateManyPerson", middleware.CreateManyPerson).Methods("POST", "OPTIONS")
	r.HandleFunc("/UpdatePerson", middleware.UpdatePerson).Methods("GET", "OPTIONS")
	r.HandleFunc("/DeletePerson", middleware.DeletePerson).Methods("GET", "OPTIONS")
	
	r.HandleFunc("/CreateManyPlaces", middleware.CreateManyPlaces).Methods("POST", "OPTIONS")
	return r
}