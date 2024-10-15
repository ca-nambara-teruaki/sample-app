package main

import (
	"context"
	"fmt"
	"log"

	"entdemo/ent"

	"github.com/caarlos0/env/v7"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string `env:"DB_HOST"     envDefault:"localhost"`
	Port     string `env:"DB_PORT"     envDefault:"5432"`
	DBName   string `env:"DB_NAME"     envDefault:"sampledb"`
	User     string `env:"DB_USER"     envDefault:"admin"`
	Password string `env:"DB_PASSWORD" envDefault:"admin"`
	SSLMode  string `env:"DB_SSLMODE"  envDefault:"disable"`
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.Create().SetAge(30).SetName("Alice").Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func main() {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("failed to parse env: %v", err)
	}

	// DSN (Data Source Name)
	dsn := "host=" + cfg.Host + " port=" + cfg.Port + " user=" + cfg.User + " dbname=" + cfg.DBName + " password=" + cfg.Password + " sslmode=" + cfg.SSLMode

	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	// Migration
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Create User Entity
	CreateUser(context.Background(), client)
}
