package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bhuvanesh/go-coffee-server/db"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

type Application struct{
	Config Config
}

func (app *Application) Serve() error {


	fmt.Println("API is listening on port: ", app.Config.Port)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", app.Config.Port),
		//TODO: add router
	}

	return srv.ListenAndServe()

}

func main(){
	err := godotenv.Load()
	if err!= nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	cfg := Config{
		Port: port,
	}

	//TODO connection to DB
	dsn := os.Getenv("DSN")
	dbConn, err := db.ConnectPostgres(dsn)

	if err!=nil {
		log.Fatal("Cannot connect to database")
	}

	defer dbConn.dbase.close()

	app := &Application{
		Config: cfg,
		//Add Models later
	}

	err = app.Serve()
	if err!= nil {
		log.Fatal(err)
	}

}