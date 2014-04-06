/*database call sql, go-sqlite3*/

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, _ := sql.Open("sqlite3", "./test.db")
	_, err := db.Exec("CREATE TABLE y(a, b, c, d)")  // внутри кавычек запрос на языке БД 
	fmt.Println(err)
	db.Close()
}