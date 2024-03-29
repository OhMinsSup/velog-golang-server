package app

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/OhMinsSup/story-server/aws"
	"github.com/OhMinsSup/story-server/ent"
	"github.com/OhMinsSup/story-server/ent/migrate"
	"github.com/OhMinsSup/story-server/libs"
	"github.com/OhMinsSup/story-server/middlewares"
	"github.com/facebook/ent/dialect"
	entsql "github.com/facebook/ent/dialect/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
)

func New() (*gin.Engine, *ent.Client) {
	// initializes database
	dbUser := libs.GetEnvWithKey("POSTGRES_USER")
	dbPassword := libs.GetEnvWithKey("POSTGRES_PASSWORD")
	dbName := libs.GetEnvWithKey("POSTGRES_DB")
	dbHost := libs.GetEnvWithKey("POSTGRES_HOST")
	dbPort := libs.GetEnvWithKey("POSTGRES_PORT")
	// https://gobyexample.com/string-formatting
	dbConfig := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPassword)

	// database tcp/ip connection error:: https://stackoverflow.com/questions/37307346/is-the-server-running-on-host-localhost-1-and-accepting-tcp-ip-connections
	db, err := sql.Open("pgx", dbConfig)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	
	drv := entsql.OpenDB(dialect.Postgres, db)
	// create database client
	client := ent.NewClient(ent.Driver(drv))

	ctx := context.Background()
	// run the auto migration tool.
	if err := client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true)); err != nil {
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
		c.Set("db", db)
		c.Set("client", client)
		c.Next()
	})

	allowOrigins := []string{"https://storeis.vercel.app"}

	if libs.GetEnvWithKey("APP_ENV") == "development" {
		allowOrigins = append(allowOrigins, "http://localhost:3000")
	}

	// setting cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = allowOrigins
	corsConfig.AllowCredentials = true
	corsConfig.AllowMethods = []string{"POST, OPTIONS, GET, PUT, PATCH, DELETE"}

	// auth middleware and set cors
	app.Use(cors.New(corsConfig))
	app.Use(middlewares.ConsumeUser(client))

	return app, client
}
