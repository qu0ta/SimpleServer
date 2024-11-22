// Package pkg contains utility functions
//
// This package includes utility functions:
// - InitDB
// - GetUserByName

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

// User is a struct for user
type User struct {
	Name     string
	Email    string
	Password string
}

// InitDB initializes the database connection pool.
//
// It parses the database configuration using the pgxpool.ParseConfig function.
// If there is an error parsing the configuration, it logs the error and exits.
// It sets the maximum number of connections, minimum number of connections,
// maximum connection lifetime, and health check period in the configuration.
//
// It creates a context and uses it to create a new database connection pool
// using the parsed configuration. If there is an error creating the pool,
// it logs the error and exits.
//
// Finally, it logs a success message indicating that the database connection
// was successfully established.
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

// GetUserByName retrieves a user from the database by their name.
//
// Parameters:
// - c: the gin context object representing the HTTP request and response.
// - name: the name of the user to retrieve.
//
// Returns:
// - A pointer to a User struct representing the retrieved user, or nil if an error occurred.
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
