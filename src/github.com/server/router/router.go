package router


import (
	"github.com/gorilla/mux"
	"github.com/server/middleware/app/users"
	"github.com/server/middleware/app/products"
	"net/http"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	r:= mux.NewRouter()
    //r.Use(commonMiddleware)
	r.HandleFunc("/", users.CreatePerson).Methods("POST", "OPTIONS")
	r.HandleFunc("/getPerson", users.GetPerson).Methods("GET", "OPTIONS")
	r.HandleFunc("/getAllPerson", users.GetAllPersons).Methods("GET", "OPTIONS")
	r.HandleFunc("/CreateManyPerson", users.CreateManyPerson).Methods("POST", "OPTIONS")
	r.HandleFunc("/UpdatePerson", users.UpdatePerson).Methods("GET", "OPTIONS")
	r.HandleFunc("/DeletePerson", users.DeletePerson).Methods("GET", "OPTIONS")
	r.HandleFunc("/GetUserPlaces", users.GetUserPlaces).Methods("GET", "OPTIONS")
	
	
	r.HandleFunc("/CreateManyPlaces", products.CreateManyPlaces).Methods("POST", "OPTIONS")
	r.HandleFunc("/GetAllPlaces", products.GetAllPlaces).Methods("GET", "OPTIONS")
	return r
}

func commonMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Content-Type", "application/json")
        next.ServeHTTP(w, r)
    })
}