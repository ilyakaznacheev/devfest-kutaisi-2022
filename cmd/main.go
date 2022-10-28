package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ilyakaznacheev/devfest-kutaisi-2022/internal/api"
	"github.com/ilyakaznacheev/devfest-kutaisi-2022/internal/database"
	"golang.org/x/sync/errgroup"
)

func main() {
	log.Println("starting application")

	dbCollection := os.Getenv("DB_COLLECTION")
	dbAddress := os.Getenv("DB_ADDRESS")
	apiAddress := os.Getenv("API_ADDRESS")

	log.Printf("connecting to database at %q", dbAddress)
	db, err := database.New(dbAddress, "", dbCollection)
	if err != nil {
		log.Fatalf("connecting to database: %v", err)
	}

	srv := api.New(apiAddress, db)

	g, _ := errgroup.WithContext(context.Background())

	log.Printf("starting server at %q", apiAddress)
	g.Go(srv.Start)

	g.Go(func() error {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		log.Println("exiting application")
		return srv.Stop()
	})

	g.Wait()
}
