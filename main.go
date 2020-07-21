package main

import (
	"context"
	"github.com/OhMinsSup/story-server/apis"
	"github.com/OhMinsSup/story-server/database"
	"github.com/OhMinsSup/story-server/middlewares"
	"github.com/OhMinsSup/story-server/storage"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func loadEnv(filename string) {
	err := godotenv.Load(filename)
	if err != nil {
		panic(err)
	}
}

func main() {
	env := os.Getenv("APP_ENV")

	switch env {
	case "production":
		loadEnv(".env.prod")
		break
	case "development":
		loadEnv(".env.dev")
		break
	default:
		loadEnv(".env.prod")
		break
	}

	// initializes database
	db, _ := database.Initialize()
	sess := storage.Initialize()

	port := os.Getenv("PORT")
	// create gin app
	app := gin.Default()
	app.MaxMultipartMemory = 8 << 20  // 8 MiB

	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(storage.Inject(sess))
	app.Use(database.Inject(db))
	app.Use(middlewares.CORSMiddleware())
	app.Use(middlewares.ConsumeUser(db))

	apis.ApplyRoutes(app)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: app,
	}

	log.Println("Listening and serving HTTP on :" + port)

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
