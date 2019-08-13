package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewMySQL(url string) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(10)
	err = db.Ping()
	return db, err
}
