package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using Railway environment")
	}

	// Read ENV
	dbURL := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// DB connect
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Ping DB
	if err = db.Ping(); err != nil {
		log.Fatal("Database unreachable:", err)
	}

	// Gin router
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "Backend live 🎉"})
	})

	r.GET("/db", func(c *gin.Context) {
		var now string
		err := db.QueryRow("SELECT NOW()").Scan(&now)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"current_time": now})
	})

	r.Run(":" + port)
}
