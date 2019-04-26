package database

import (
	"database/sql"
	"errors"
	//"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var MaxConnections = 3

const IdleTimeMinutes int = 10

func initializeDB(connectionURL string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connectionURL)

	if err != nil {
		log.Printf("mysql.initializeDB - Erro ao abrir conexao com DB: %v", err)
		return nil, err
	}

	db.SetMaxIdleConns(MaxConnections) // Número máximo de conexões abertas sem uso
	db.SetMaxOpenConns(MaxConnections) // Número máximo de conexões abertas simultaneamente

	var connectionTimeout = time.Minute * time.Duration(IdleTimeMinutes)
	db.SetConnMaxLifetime(connectionTimeout) // Tempo máximo que uma conexão pode ser reutilizada. 0 = para sempre

	err = connectionValidate(db)

	if err != nil {
		db = nil
		log.Printf("mysql.initializeDB - Falha ao testar conexao com DB.")
		return nil, err
	}

	return db, nil
}

func connectionValidate(db *sql.DB) error {

	if db == nil {
		return errors.New("mysql.connectionValidate - Ponteiro para conexão com o DB nulo")
	}

	err := db.Ping()

	if err != nil {
		log.Println("mysql.connectionValidate - Erro testar conexao com DB.")
		return err
	}

	return nil
}


func closeDB(db *sql.DB) error {
	if db == nil {
		return errors.New("mysql.closeDB - Ponteiro para conexão com o DB nulo")
	}
	err := db.Close()
	db = nil

	if err != nil {
		log.Println("mysql.closeDB - Erro ao finalizar banco da dados.")
		return err
	}

	return nil
}
