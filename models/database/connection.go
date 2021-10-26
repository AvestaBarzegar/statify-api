package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/AvestaBarzegar/statify-api/helpers/consts"
	_ "github.com/lib/pq"
)

type Database struct {
	Conn *sql.DB
}

func Initialize() (Database, error) {
	db := Database{}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		consts.HOST, consts.GetPortDB(), consts.USER, consts.PASSWORD, consts.DB)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return db, err
	}
	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}
	log.Println("Database connection established")
	return db, nil
}
