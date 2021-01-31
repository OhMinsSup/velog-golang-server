package app

import (
	"context"
	"fmt"
	"github.com/OhMinsSup/story-server/ent"
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/OhMinsSup/story-server/helpers/aws"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func New() *gin.Engine {
	allowOrigins := []string{"https://storeis.vercel.app"}

	if helpers.GetEnvWithKey("APP_ENV") == "development" {
		allowOrigins = append(allowOrigins, "http://localhost:3000")
	}

	// initializes database
	//db, _ := database.Initialize()

	dbUser := helpers.GetEnvWithKey("POSTGRES_USER")
	dbPassword := helpers.GetEnvWithKey("POSTGRES_PASSWORD")
	dbName := helpers.GetEnvWithKey("POSTGRES_DB")
	dbHost := helpers.GetEnvWithKey("POSTGRES_HOST")
	dbPort := helpers.GetEnvWithKey("POSTGRES_PORT")
	// https://gobyexample.com/string-formatting
	dbConfig := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPassword)

	// create database client
	client, err := ent.Open("postgres", dbConfig)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)

	}

	// database close
	defer client.Close()
	ctx := context.Background()
	// run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// initializes aws s3 storage
	sess := aws.Initialize()

	// create gin app
	app := gin.Default()
	app.MaxMultipartMemory = 8 << 20 // 8 MiB

	// setting middleware
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(aws.Inject(sess))
	app.Use(func(c *gin.Context) {
		c.Set("client", client)
		c.Next()
	})
	//app.Use(database.Inject(db))

	// setting cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = allowOrigins
	corsConfig.AllowCredentials = true
	corsConfig.AllowMethods = []string{"POST, OPTIONS, GET, PUT, PATCH, DELETE"}

	// auth middleware and set cors
	app.Use(cors.New(corsConfig))
	//app.Use(middlewares.ConsumeUser(db))

	return app
}
