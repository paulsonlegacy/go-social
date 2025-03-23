package db

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

// NewDBConnection initializes a new DB connection pool
func NewDBConnection(dburl string, maxOpenConns, maxIdleConns int, maxIdleTime string) (*sql.DB, error) {
	// Converting maxIdleTime to time object
	duration, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		return nil, err
	}
	
	// Making a postgres DB connection
	db, err := sql.Open("mysql", dburl)
	if err != nil {
		return nil, err
	} else {
		log.Println("DB connection successful!")
	}

	// Configuration DB
	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxIdleTime(duration)
	db.SetMaxIdleConns(maxIdleConns)

	// Creating a timeout context for DB connections
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}