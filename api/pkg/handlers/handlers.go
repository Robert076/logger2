package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Robert076/logger2.git/api/pkg/message"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

func InitDB() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
	dbConfig := DBConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
		Port:     os.Getenv("POSTGRES_PORT"),
	}
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("error opening db: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging db: %v", err)
	}

	return db, nil
}

func HandlerPost(context *gin.Context) {
	db, err := InitDB()
	if err != nil {
		http.Error(context.Writer, "Could not initialize DB", http.StatusInternalServerError)
		log.Printf("Could not open db: %v", err)
		return
	}

	var msg message.Message

	if err := json.NewDecoder(context.Request.Body).Decode(&msg); err != nil {
		http.Error(context.Writer, "Could not read request body", http.StatusBadRequest)
		log.Printf("Could not read request body: %v", err)
		return
	}

	insertQuery := `INSERT INTO messages(message, created_at) VALUES($1, $2)`
	if _, err = db.Exec(insertQuery, msg.Message, time.Now()); err != nil {
		http.Error(context.Writer, "Could not insert into DB", http.StatusBadRequest)
		log.Printf("Could not insert into DB: %v", err)
		return
	}
}

func HandlerGet(context *gin.Context) {
	db, err := InitDB()

	if err != nil {
		http.Error(context.Writer, "Could not initialize DB", http.StatusInternalServerError)
		log.Printf("Could not open db: %v", err)
		return
	}

	rows, err := db.Query(`SELECT * FROM "messages"`)
	if err != nil {
		http.Error(context.Writer, "Could not query DB", http.StatusInternalServerError)
		log.Printf("Could not query db: %v", err)
		return
	}

	defer rows.Close()

	var msg message.Message
	for rows.Next() {
		err := rows.Scan(&msg.Id, &msg.Message, &msg.CreatedAt)
		if err != nil {
			http.Error(context.Writer, "Could not scan DB row", http.StatusInternalServerError)
			log.Printf("Could not scan db row: %v", err)
			return
		}

		enc := json.NewEncoder(context.Writer)
		enc.SetIndent("", "    ")

		if err := enc.Encode(msg); err != nil {
			http.Error(context.Writer, "Could not encode message", http.StatusInternalServerError)
			log.Printf("Could not encode message: %v", err)
			return
		}
	}

	err = rows.Err()
	if err != nil {
		http.Error(context.Writer, "Error occured when reading db", http.StatusInternalServerError)
		log.Printf("Error occured when reading db: %v", err)
		return
	}
}
