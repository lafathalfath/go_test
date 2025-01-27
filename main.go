package main

import (
	"database/sql"
	"fmt"
	"go_test/controller"
	"go_test/internal/database"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type ApiConfig struct {
	DB *database.Queries
}

func main() {
	// DB CONNECTION
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	db := database.New(conn)
	// END DB CONNECTION
	
	
	// ROUTER
	apiConfig := controller.ApiConfig{DB: db}
	r := mux.NewRouter()
    r.HandleFunc("/products", apiConfig.GetAllProducts).Methods("GET")
    r.HandleFunc("/products", apiConfig.CreateProduct).Methods("POST")
    r.HandleFunc("/products/{id}", apiConfig.UpdateProduct).Methods("PUT")
    r.HandleFunc("/products/{id}", apiConfig.DeleteProduct).Methods("DELETE")
	// END ROUTER
    
	// APLICATION LISTENER
	log.Println("server is running on port: 8080")
	http.ListenAndServe(":8080", r)
	// END APLICATION LISTENER
}