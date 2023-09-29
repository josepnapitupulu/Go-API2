package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/josepnapitupulu/API_Create_Nikah/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterSidisRoutes(r)
	http.Handle("/", r)

	fmt.Print("Starting Server localhost:7072")
	log.Fatal(http.ListenAndServe("localhost:7072", r))
}