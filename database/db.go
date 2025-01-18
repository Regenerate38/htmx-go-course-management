package database

import (
	"database/sql"
	"fmt"
	"htmx-go-course-management/config"

	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	host := config.Config("DB_HOST")
	portStr := config.Config("DB_PORT")
	user := config.Config("DB_USER")
	password := config.Config("DB_PASSWORD")
	dbName := config.Config("DB_NAME")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid port number: %s. Error: %v", portStr, err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, password, host, port, dbName)
	fmt.Println(dsn)
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	log.Println("Database connection established successfully!")
}
