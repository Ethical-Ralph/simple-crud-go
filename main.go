package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Ethical-Ralph/simple-crud-go/database"
	"github.com/Ethical-Ralph/simple-crud-go/handlers"
	"github.com/joho/godotenv"
)

// ex:product data
// {
//     "name": "Milk",
//     "description": "A milk",
//     "price": 1000
// }

// Create Product "/"
// Method: POST

// Get all Product "/"
// Method: GET

// Get a Product "/:id"
// Method: GET

// Update a Product "/:id"
// Method: PUT

// Delete a Product "/:id"
// Method: DELETE


func main() {
	dbUri := getEnv("MONGODB_URI")
	port := ":3000"

	if dbUri == "" {
		log.Fatal("mongodb uri not found")
	}

	db := database.NewDatabase(dbUri)
	l := log.New(os.Stdout, "api", log.LstdFlags)
	pd := handlers.Product(l, db)

	sm := http.NewServeMux()
	sm.Handle("/", pd)

	s:= &http.Server{
		Addr: port,
		Handler: sm,
	}

	done := make(chan bool)
	go s.ListenAndServe()
	log.Printf("App started on port %s", port)
	<- done
}

func getEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Count't load env file")
	}

	return os.Getenv(key)
}