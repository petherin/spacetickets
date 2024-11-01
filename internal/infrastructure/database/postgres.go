package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/petherin/spacetickets/internal/infrastructure/config"
)

const databaseDriver = "postgres"

// PostGres encapsulates objects needed to interact with a PostGres database.
type PostGres struct {
	Repo *sql.DB
}

// New returns a new PostGres.
func New(cfg config.Config) (*PostGres, error) {
	db, err := connect(cfg)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.DBMaxOpenConns)
	db.SetMaxIdleConns(cfg.DBMaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(cfg.DBConnMaxLifetimeSecs) * time.Second)

	return &PostGres{Repo: db}, nil
}

// Close closes the connection to the database.
func (p *PostGres) Close() {
	p.Repo.Close()
}

func connect(cfg config.Config) (*sql.DB, error) {
	maxRetries := cfg.DBConnRetries
	retryInterval := time.Duration(cfg.DBConnRetryIntervalSecs) * time.Second
	var db *sql.DB
	var err error

	log.Println("Connecting to database...")

	for i := 0; i < maxRetries; i++ {
		connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s", cfg.DBUsername, cfg.DBPassword, cfg.DBName, cfg.DBHost)
		db, err = sql.Open(databaseDriver, connectionString)
		if err == nil {
			log.Println("Connected to database")
			break
		}

		log.Printf("Failed to connect to the database (attempt %d/%d): %v\n", i+1, maxRetries, err)
		time.Sleep(retryInterval)
	}

	log.Println("Pinging database...")
	pinged := false
	for i := 0; i < maxRetries; i++ {
		err = db.Ping()
		if err == nil {
			log.Println("Pinged database")
			pinged = true
			break
		}

		log.Printf("Failed to ping database (attempt %d/%d): %v\n", i+1, maxRetries, err)
		time.Sleep(retryInterval)
	}

	if !pinged {
		return nil, fmt.Errorf("failed to ping database")
	}

	return db, nil
}
