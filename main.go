package main

import (
	"fmt"
	"log"
	"time"

	"github.com/SteakBarbare/Partiel-MT4/database"
	"github.com/SteakBarbare/Partiel-MT4/routes"
	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func main() {

	time.Sleep(5 * time.Second)

	cfg := database.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg)

	database.Connect(cfg)
	database.MakeMigrations()

	r := gin.Default()
	
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "Authorization",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	routes.InitializeRoutes(r)

	log.Fatal(r.Run(":8080")) // listen and serve on 8080
}
