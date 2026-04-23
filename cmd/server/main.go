package main

import (
	"backend/internal/database"
	"backend/internal/middleware"
	"backend/internal/routes"
	"log"
	"net/http"
)

func main(){
	db, dbErr  := database.ConnectDB()
	
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	defer db.Close()

	initErr := database.InitializeDB(db)

	if initErr != nil {
		log.Fatal(initErr)
	}

	seedErr := database.SeedDatabase(db)
	if seedErr != nil {
		log.Fatal(seedErr)
	}
	
	mux := http.NewServeMux()

	routes.RegisterRoutes(mux, db)

	handler:= middleware.EnableCors(mux)

	err := http.ListenAndServe(":8080", handler)
	
	if err != nil {
		log.Fatal(err)
	}
}