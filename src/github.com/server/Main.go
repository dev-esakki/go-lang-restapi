package main

import (    
    "log"
    "net/http"
    "github.com/server/router"
)


func main() {
    r := router.Router()    
    const port string = ":8000"
    log.Println("server listening on ", port)
    log.Fatalln(http.ListenAndServe(port, r))
}