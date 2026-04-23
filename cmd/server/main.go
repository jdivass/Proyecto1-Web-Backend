package main

import (
	"backend/internal/database"
	"backend/internal/middleware"
	"backend/internal/routes"
	"log"
	"net/http"
	"backend/internal/utils"
)

func main(){
	mux := http.NewServeMux()

	if err := utils.EnsureUploadsDir(); err != nil {
	log.Fatal(err)
	}
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

	routes.RegisterRoutes(mux, db)

	fileServer := http.FileServer(http.Dir("./uploads"))
	mux.Handle("/uploads/", http.StripPrefix("/uploads/", fileServer))

	handler:= middleware.EnableCors(mux)

	err := http.ListenAndServe(":8080", handler)
	log.Println("server running on :8080")
	if err != nil {
		log.Fatal(err)
	}
}