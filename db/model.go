package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/miraklik/TODO-list/config"
)

func ConnectDB() (*pgx.Conn, error) {
	cfg, err := config.Load()
	if err != nil {
		log.Printf("Failed to get .env: %v", err)
		return nil, err
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Database.Db_host, cfg.Database.Db_port, cfg.Database.Db_user, cfg.Database.Db_pass, cfg.Database.Db_name)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Printf("Failed to connect db: %v", err)
		return nil, nil
	}

	if err := db.Ping(ctx); err != nil {
		log.Printf("Failed to ping database: %v", err)
		db.Close(ctx)
		return nil, err
	}

	log.Println("Successfully connected to database")
	return db, nil
}

func InitSchema() error {
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT,
		status TEXT CHECK (status IN ('new', 'in_progress', 'done')) DEFAULT 'new',
		created_at TIMESTAMP DEFAULT now(),
		updated_at TIMESTAMP DEFAULT now()
	);
	`

	db, _ := ConnectDB()

	_, err := db.Exec(context.Background(), query)
	if err != nil {
		log.Printf("schema init error: %v\n", err)
	}
	return err
}
