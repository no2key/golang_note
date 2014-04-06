/*add new data in database*/

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, _ := sql.Open("sqlite3", "./test.db")
	tx, err_4 := db.Begin()
	stmt, err_2 := tx.Prepare("insert into y(a, d) values(?, ?)")

	for i := 0; i < 8; i++ {
		_, err_3 := stmt.Exec(i, "text")
		fmt.Println(err_3)
	}

	fmt.Println(err_2)
	fmt.Println(err_4)

	stmt.Close()
	tx.Commit()
	db.Close()
}

// возможно много лишних комманд
например зачем begin или зачем commit