package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	pq "github.com/lib/pq"
)

func main() {
	r := gin.Default()
	// static content populates within docker container
	r.Use(static.Serve("/", static.LocalFile("./web", true)))
	api := r.Group("/api")
	dbURL := os.Getenv("DATABASE_URL")
	log.Printf("DB [%s]", dbURL)
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("error opening db: %q", err)
	}

	log.Println("looking good!")
	api.GET("/ping", pingFunc(db))

	r.Run()
}

func registerPing(db *sql.DB) {
	_, err := db.Exec("INSERT INTO ping_timestamp (occurred) VALUES ($1)", time.Now())
	if err != nil {
		log.Println("Unable to insert ping")
		log.Println(err.Error())
	}
}

func pingFunc(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer registerPing(db)
		r := db.QueryRow("SELECT occurred FROM ping_timestamp ORDER BY id DESC LIMIT 1")
		var lastDate pq.NullTime
		r.Scan(&lastDate)

		message := "hit refresh to see how long since the last ping!"
		if lastDate.Valid {
			message = fmt.Sprintf("%v ago", time.Now().Sub(lastDate.Time).String())
		}

		c.JSON(200, gin.H{
			"message": message,
		})
	}
}
