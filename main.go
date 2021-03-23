package main

import (
	"context"
	"encoding/gob"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/database"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/domain"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/models"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/rabbitmq"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	port := "8080"
	newRouter := router.NewRouter()

	err := database.ConnectSQL()
	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}

	err = rabbitmq.ConnectToRabbitMq()
	if err != nil {
		log.Fatalf("could not connect to rabbitMQ: %v", err)
	}

	ctx := context.Background()

	foreverLoopDelay := 5 * time.Second

	if _, err := database.GetEsConn(ctx, foreverLoopDelay); err != nil {
		log.Fatalf("could not connect to es: %v", err)
	}

	log.Printf("database.EsConn: %v", database.EsConn)

	domain.InitBusses()
	gob.Register(models.Order{})
	gob.Register(models.OrderLine{})

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
