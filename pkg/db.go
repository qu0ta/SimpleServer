package pkg

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"log/slog"
	"time"
)

var Pool *pgxpool.Pool

func InitDB() {
	cfg, err := pgxpool.ParseConfig("postgres://postgres:gupaljuk@localhost:1111/postgres")
	if err != nil {
		slog.Error("Error while parsing DB config: " + err.Error())
		log.Fatalln(err.Error())
		return
	}
	cfg.MaxConns = 10
	cfg.MinConns = 5
	cfg.MaxConnLifetime = 5 * time.Minute
	cfg.HealthCheckPeriod = 1 * time.Minute

	ctx := context.Background()

	Pool, err = pgxpool.New(ctx, cfg.ConnString())
	if err != nil {
		slog.Error("Error while initialize DB: " + err.Error())
		log.Fatalln(err.Error())
		return
	}
	log.Println("Successfully connected to DB!")
}
