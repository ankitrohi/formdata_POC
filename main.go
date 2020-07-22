package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

var conn *sql.DB

func init() {
	initDB()
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.Use(
		handlers.RecoveryHandler(),
		handlers.CompressHandler,
	)
	router.HandleFunc("/imports/", ReadAndStore).Methods("POST")
	allowedHeaders := []string{"X-Requested-With", "Content-Type", "Authorization"}
	allowedMethods := []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}
	allowedOrigins := []string{"*"}
	http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedHeaders(allowedHeaders), handlers.AllowedMethods(allowedMethods), handlers.AllowedOrigins(allowedOrigins))(handlers.LoggingHandler(os.Stdout, router)))
}

func initDB() {
	var err error
	conn, err = sql.Open("mysql", "root:root@tcp(mysql:3306)/test_db")
	if err != nil {
		fmt.Printf("error in initialising database: %v", err)
		os.Exit(1)
	}

	err = conn.Ping()
	if err != nil {
		fmt.Printf("error in connecting to database: %v", err)
		os.Exit(1)
	}

	conn.SetMaxOpenConns(100)                 // Set maximum open connections 100
	conn.SetMaxIdleConns(25)                  // Set maximum idle connection 0
	conn.SetConnMaxLifetime(30 * time.Minute) // Set maximum connection life time

	_, err = conn.Query("CREATE TABLE info(name VARCHAR(25), email VARCHAR(100))")
	if err != nil {
		fmt.Printf("error in creating table: %v", err)
		os.Exit(1)
	}
}

func ReadAndStore(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(32 << 20)

	name := r.FormValue("name")
	email := r.FormValue("email")

	query := "INSERT INTO info(name, email) VALUES(?,?)"

	_, err := conn.Query(query, name, email)
	if err != nil {
		fmt.Fprintf(w, "error in insertin data: " + err.Error())
		return
	}

	fmt.Fprintf(w, "successfully added the data to database")
}