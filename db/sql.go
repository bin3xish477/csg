package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DB *sql.DB
)

func CreateDB(csgFolder, vault string) {
	f, err := os.Create(vault)
	if err != nil {
		log.Printf("unable to create directory in folder: %s", csgFolder)
	}
	defer f.Close()

	err = os.Chmod(vault, 0755)
	if err != nil {
		panic(err)
	}

	stmt := "CREATE TABLE `vault` ( `uid` INTEGER PRIMARY KEY AUTOINCREMENT, `host` VARCHAR(64) NULL, `app` VARCHAR(64) NULL, `user` VARCHAR(64) NULL, `passwd` VARCHAR(64) NULL);"
	fmt.Printf("executing \n%s\n", stmt)
	_, err = DB.Exec(stmt)
	if err != nil {
		log.Fatalln("unable to create database")
	}
}
