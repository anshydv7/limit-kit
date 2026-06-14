package database

import (
	"auth-service/base/config"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitiateDatabase() error {
	uri := config.Config.Database.Uri

	var err error
	db, err = sql.Open("postgres", uri)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// 1. Database connection pool configuration
	maxOpen := config.Config.Database.MaxOpenConns
	if maxOpen <= 0 {
		maxOpen = 25 // default fallback
	}
	db.SetMaxOpenConns(maxOpen)

	maxIdle := config.Config.Database.MaxIdleConns
	if maxIdle <= 0 {
		maxIdle = 25 // default fallback
	}
	db.SetMaxIdleConns(maxIdle)

	lifetimeMins := config.Config.Database.ConnMaxLifetimeMins
	if lifetimeMins <= 0 {
		lifetimeMins = 5 // default fallback
	}
	db.SetConnMaxLifetime(time.Duration(lifetimeMins) * time.Minute)

	// 2. Connection timeout for the initial ping
	timeoutSecs := config.Config.Database.ConnectionTimeoutSeconds
	if timeoutSecs <= 0 {
		timeoutSecs = 10 // default fallback
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSecs)*time.Second)
	defer cancel()

	// Ping the database with context to enforce the timeout
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Successfully connected to the database")
	return nil
}

func GetDB() *sql.DB {
	return db
}
