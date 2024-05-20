package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/vedantwankhade/go-microservices/logger-service/data"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	port     = ":80"
	rpcPort  = ":5001"
	mongoURL = "mongodb://mongo:27017"
	grpcPort = ":50001"
)

var client *mongo.Client

type Config struct {
	Models data.Models
}

func main() {
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}
	client = mongoClient
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	app := Config{Models: data.New(client)}
	// go app.serve()
	log.Println("Starting logging server on", port)
	srv := &http.Server{
		Addr:    port,
		Handler: app.routes(),
	}
	log.Fatal(srv.ListenAndServe())
}

func (app *Config) serve() {
	srv := &http.Server{
		Addr:    port,
		Handler: app.routes(),
	}
	log.Fatal(srv.ListenAndServe())
}

func connectToMongo() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("error connecting:", err)
		return nil, err
	}
	return c, nil
}