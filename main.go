package main

import (
	"github.com/HETIC-MT-P2021/gocqrs/database"
	"github.com/HETIC-MT-P2021/gocqrs/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	port := "8080"
	newRouter := router.NewRouter()

	err := database.ConnectSQL()
	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}

	esCfg := &database.ConfigEs{URL: "http://es:9200"}

	if err := database.ConnectES(esCfg); err != nil {
		log.Fatalf("could not connect to es: %v", err)
	}

	go func() {
		log.Print("\nServer started on port " + port)
		if err := http.ListenAndServe(":"+port, newRouter); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error : %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

}
