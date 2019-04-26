package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
)

type ServerConfig struct {
	DB *sql.DB
}

var dbConfig *ServerConfig

func StartDB() error {
	var err error

	dbConfig = new(ServerConfig)

	//connectionString := ""fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", "root", "root", "localhost:3306", "school")
	connectionString := os.Getenv("JAWSDB_MARIA_URL")

	dbConfig.DB, err = initializeDB(connectionString)

	if err != nil {
		log.Println("db.InicializaDB - Erro ao inicializar banco de dados.")

		return err
	}

	if dbConfig.DB == nil {
		log.Panicln("db.InicializaDB - Ponteiro para conexão com DB nulo.")
		return nil
	}
	return nil
}

func Finalize() {
	closeDB(dbConfig.DB)
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := dbConfig.DB.Prepare(query)

	if err != nil {
		log.Printf("db.Query - Erro ao preparar a query: " + query)
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		log.Printf("db.Query - Erro ao executar a query: " + query)
		return nil, err
	}

	if rows == nil {
		log.Println("db.Query - rows nulo.")
		return nil, errors.New("db.Query - rows nulo")
	}

	return rows, nil
}

func Exec(query string, args ...interface{}) (*sql.Result, error) {
	stmt, err := dbConfig.DB.Prepare(query)

	if err != nil {
		log.Println("Exec - Erro ao preparar a query: " + query)
		log.Printf("Exec - Erro: %v", err)
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(args...)
	if err != nil {
		log.Printf("Exec - Erro ao executar a query: " + query)
		fmt.Println(err.Error())
		return nil, err
	}

	return &result, nil
}
