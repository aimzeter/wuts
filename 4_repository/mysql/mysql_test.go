package mysql_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aimzeter/wuts/4_repository/mysql"
	"github.com/jmoiron/sqlx"
)

func setupDB(t *testing.T) *sqlx.DB {
	t.Helper()
	url := "root@(localhost:3306)/zone_test?parseTime=true"

	db, err := mysql.NewMySQL(url)
	if err != nil {
		t.Fatalf("Fail to connect to DB, got error %s", err.Error())
	}
	return db
}

func setupMockDB(t *testing.T) (*sqlx.DB, sqlmock.Sqlmock, func() error) {
	t.Helper()
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error '%s' when opening a stub database connection", err)
	}

	db := sqlx.NewDb(mockDB, "mysql")
	return db, mock, mockDB.Close
}
