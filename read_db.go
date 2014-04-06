/*database read + other*/

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func main() {
	os.Remove("./test.db")

	db, _ := sql.Open("sqlite3", "./test.db")
	_, err := db.Exec("CREATE TABLE x(a, d)")
	tx, err_4 := db.Begin()
	stmt, err_2 := tx.Prepare("insert into x(a, d) values(?, ?)")

	for i := 0; i < 8; i++ {
		_, err_3 := stmt.Exec(i, "test")
		fmt.Println(err_3)
	}

	fmt.Println(err)
	fmt.Println(err_2)
	fmt.Println(err_4)

	stmt.Close()
	tx.Commit()

	rows, err_5 := db.Query("select a, d from x")
	fmt.Println(err_5)

	for rows.Next() {
		var t1 int
		var t2 string
		rows.Scan(&t1, &t2)
		fmt.Println(t1, t2)
	}

rows.Close()
	db.Close()
}

// возможны лишии команды