package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"

	"gopkg.in/gorp.v2"
)

var dbm *gorp.DbMap

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot load .env file")
		os.Exit(1)
	}

	addr := os.Getenv("PORT")
	datasource := os.Getenv("DB_DATASOURCE")

	db, err := sql.Open("mysql", datasource)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	dbm = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	defer dbm.Db.Close()

	dbm.AddTableWithName(User{}, "user")

	r := mux.NewRouter()
	r.Methods(http.MethodPost).Path("/signup").HandlerFunc(signup)
	r.Methods(http.MethodPost).Path("/login").HandlerFunc(login)

	log.SetFlags(log.Ldate + log.Ltime + log.Lshortfile)
	log.SetOutput(os.Stdout)
	log.Printf("Listening on port %s", addr)

	err = http.ListenAndServe(fmt.Sprintf(":%s", addr), handlers.CombinedLoggingHandler(os.Stdout, r))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
