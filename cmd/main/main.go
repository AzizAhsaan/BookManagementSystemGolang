package main

import(
	"log"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/AzizAhsaan/BookStoreManagementSystemGolang/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/",r)
	fmt.Println("Starting the application...")

	log.Fatal(http.ListenAndServe("localhost:9010",r))
	fmt.Println("Connected to the database.")

}