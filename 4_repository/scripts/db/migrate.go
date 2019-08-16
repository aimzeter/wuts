package main

import (
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	participantsTable = "CREATE TABLE participants (" +
		"`id` INT(11) NOT NULL AUTO_INCREMENT," +
		"`nik` VARCHAR(255)," +
		"`name` VARCHAR(255)," +
		"`address` VARCHAR(255)," +
		"`latitude` FLOAT," +
		"`longitude` FLOAT," +
		"`auto_reg` TINYINT(1) DEFAULT 0," +
		"`distance` FLOAT," +
		"`total_score` FLOAT," +
		"PRIMARY KEY (`id`)" +
		") ENGINE = InnoDB DEFAULT CHARSET = utf8"

	studentsTable = "CREATE TABLE students (" +
		"`id` INT(11) NOT NULL AUTO_INCREMENT," +
		"`nik` VARCHAR(255)," +
		"`name` VARCHAR(255)," +
		"`address` VARCHAR(255)," +
		"`latitude` FLOAT," +
		"`longitude` FLOAT," +
		"PRIMARY KEY (`id`)" +
		") ENGINE = InnoDB DEFAULT CHARSET = utf8"

	dbName = "zone_test"
)

func main() {
	host := flag.String("host", "127.0.0.1", "your mysql host")
	port := flag.String("port", "3306", "your mysql port")
	username := flag.String("username", "", "your mysql username")
	password := flag.String("password", "", "your mysql password")
	flag.Parse()

	dataSourceName := fmt.Sprintf("%s:%s@(%s:%s)/?parseTime=true", *username, *password, *host, *port)
	fmt.Println("Connecting to", dataSourceName)

	db, _ := sqlx.Open("mysql", dataSourceName)
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(err.Error())
	}

	createDB := fmt.Sprintf("CREATE DATABASE %s", dbName)
	db.Exec(createDB)
	fmt.Println(createDB)

	useDB := fmt.Sprintf("USE %s", dbName)
	db.Exec(useDB)
	fmt.Println(useDB)

	fmt.Println("DROP ALL TABLE...")
	db.Exec("DROP TABLE participants")
	db.Exec("DROP TABLE students")

	fmt.Println("MIGRATING...")

	fmt.Println("CREATING TABLE participants")
	checkErr(db.Exec(participantsTable))

	fmt.Println("CREATING TABLE students")
	checkErr(db.Exec(studentsTable))

	fmt.Println("DONE")
}

func checkErr(_ sql.Result, err error) {
	if err != nil {
		fmt.Println("Got error:", err.Error())
	}
}
