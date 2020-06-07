package main

import (
	"github.com/OhMinsSup/story-server/apis"
	"github.com/OhMinsSup/story-server/database"
	"github.com/OhMinsSup/story-server/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// initializes database
	db, _ := database.Initialize()

	port := os.Getenv("PORT")
	// create gin app
	app := gin.Default()

	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(database.Inject(db))
	app.Use(middlewares.ConsumeUser(db))

	apis.ApplyRoutes(app)

	app.Run(":" + port)
}
