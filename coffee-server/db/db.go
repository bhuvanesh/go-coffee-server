package db

import (
	"database/sql"
	"fmt"
	"time"
)

type DB struct{
	dbase *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn =10
const maxIdleDbConn = 5
const maxDbLifeTime = 5 * time.Minute

func ConnectPostgres(dsn string) (*DB, error) {
	//dsn -> data source name
	d, err := sql.Open("pgx", dsn)
	if err!= nil {
		return nil, err
	}

	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxIdleTime(maxDbLifeTime)

	err = testDB(d)
	if err!= nil{
		return nil, err
	}

	dbConn.dbase = d
	return dbConn, nil
}

func testDB(d *sql.DB) error {
	err := d.Ping()
	if err!= nil{
		fmt.Println("Error: ", err)
		return err
	}

	fmt.Println("*** Pinged database successfully ***")
	return nil
}