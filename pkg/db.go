package pkg

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"log/slog"
	"net/http"
	"time"
)

var Pool *pgxpool.Pool

type User struct {
	Name     string
	Email    string
	Password string
}

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

func GetUserByName(c *gin.Context, name string) *User {
	conn, err := Pool.Acquire(c)
	if err != nil {
		BaseErrorHandler(c, err, "Error while taking connection")
		return nil
	}

	var user User
	fmt.Println(name)
	row := conn.QueryRow(c, "SELECT name, email, password FROM users WHERE name = $1", name)
	err = row.Scan(&user.Name, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusBadRequest, "Invalid credentials")
			return nil
		}
		BaseErrorHandler(c, err, "Error scanning user")
		return nil
	}
	return &user

}
