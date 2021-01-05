package main

import (
	"gocqrs/database"
	"gocqrs/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	port := "8080"
	router := router.NewRouter()

	err := database.Connect()
	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}

	go func() {
		log.Print("\nServer started on port " + port)
		if err := http.ListenAndServe(":"+port, (router)); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error : %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	if err != nil {
		log.Fatalf("could not serve on port %s", port)
	}
}
