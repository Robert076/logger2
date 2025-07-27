package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Robert076/logger2.git/api/internal/message"
	"github.com/gin-gonic/gin"
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
	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DBNAME"),
		Port:     os.Getenv("DBPORT"),
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

	insertQuery := `INSERT INTO messages(text, created_at) VALUES($1, $2)`
	if _, err = db.Exec(insertQuery, msg.Text, time.Now()); err != nil {
		http.Error(context.Writer, "Could not insert into DB", http.StatusBadRequest)
		log.Printf("Could not insert into DB: %v", err)
		return
	}
}
