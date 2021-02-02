package main

import (
	"context"
	"github.com/OhMinsSup/story-server/apis"
	"github.com/OhMinsSup/story-server/app"
	"github.com/OhMinsSup/story-server/helpers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	app.NewEnv()
}

func main() {
	port := helpers.GetEnvWithKey("PORT")
	server, client := app.New()

	apis.ApplyRoutes(server)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: server,
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

		log.Println("Database Close...")
		client.Close()
	}
	log.Println("Server exiting")
}
