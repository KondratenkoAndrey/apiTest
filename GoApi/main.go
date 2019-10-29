package main

import (
	"./controllers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/people", controllers.GetPeople).Methods("GET")

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Print(err)
	}
}
