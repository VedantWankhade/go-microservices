package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/vedantwankhade/go-microservices/auth-service/data"
)

const port = ":80"

var count int

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting auth service on", port)

	conn := connectToDB()
	if conn == nil {
		log.Panic("Cant connec to postgress")
	}

	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}
	srv := &http.Server{
		Addr:    port,
		Handler: app.routes(),
	}
	log.Panic(srv.ListenAndServe())
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")
	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgress not yet up...")
			count++
		} else {
			log.Println("Connected to postgress!")
			return connection
		}
		if count > 10 {
			log.Println(err)
			return nil
		}
		log.Println("Trying again after 2 seconds...")
		time.Sleep(time.Second * 2)
		continue
	}
}
