package database

import (
	"database/sql"
	"fmt"

	"github.com/AvestaBarzegar/statify-api/helpers/consts"
)

func OpenConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", consts.HOST, consts.GetPortDB(), consts.USER, consts.PASSWORD, consts.DB)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
