package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
)

type Asdf struct {
	Db *sql.DB
}

var (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "example"
	dbname   = "postgres"
)

func InitProdEnvironment() {
	host = os.Getenv("DB_HOST")         //"localhost"
	port = os.Getenv("DB_PORT")         // "5432"
	user = os.Getenv("DB_USERNAME")     //"postgres"
	password = os.Getenv("DB_PASSWORD") //"example"
	dbname = os.Getenv("DB_NAME")       //"postgres"
}

func SetupDatabase() (*sql.DB, error) {

	p, _ := strconv.Atoi(port)
	psqlInfo := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
		host, p, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
