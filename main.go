package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"learn-gorm/app"
	"learn-gorm/config"
)

func main() {

	config.InitConfig()

	mysqlConn, err := config.ConnectMySQL()
	if err != nil {
		log.Fatalln("Error DB connection: ", err.Error())
	}
	defer mysqlConn.Close()

	router := app.InitRouter(mysqlConn)
	log.Println("routes Initialized")

	port := config.Getenv("PORT", "8260")
	srv := &http.Server{
		Addr:    ":" + port, // PORT for calling Request
		Handler: router,     // Services Request List
	}

	log.Println("Server Initialized")
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 2 seconds.
	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")

}
