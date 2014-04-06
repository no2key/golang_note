/*check name and password in data base with gorp*/


// test_4 project main.go
package main

import (
	"database/sql"
	"fmt"
	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Name     string
	Password string
	Id       int
}

func getUser(username string, Dbm *gorp.DbMap) *User {
	users, err_3 := Dbm.Select(User{}, "SELECT * FROM User WHERE Name = ?", username)
	fmt.Println(err_3)
	fmt.Println(users[0].(*User))
	return users[0].(*User)
}

func main() {
	fmt.Println("Hello World!")
	db, err_1 := sql.Open("sqlite3", "./my.db")
	fmt.Println(err_1)

	Dbm := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	Dbm.AddTable(User{}).SetKeys(true, "Id")
	Dbm.CreateTables()

	demo_user := &User{"demo", "demo", 0}
	err_2 := Dbm.Insert(demo_user)
	fmt.Println(err_2)

	var enter_name, enter_password, enter_cmd string

	for i := 0; i < 2; i++ {
		fmt.Printf("Enter 'login' or 'create' \n")
		fmt.Printf("cmd.")
		fmt.Scanf("%s \n", &enter_cmd)

		if enter_cmd == "login" {
			fmt.Printf("Enter Name: ")
			fmt.Scanf("%s \n", &enter_name)
			fmt.Printf("Enter Password: ")
			fmt.Scanf("%s \n", &enter_password)

			user_my := getUser(enter_name, Dbm)
			if user_my.Password == enter_password {
				fmt.Printf("Password is True")
			} else {
				fmt.Printf("Password is False")
			}
		} else {

			fmt.Printf("Enter Name: ")
			fmt.Scanf("%s \n", &enter_name)
			fmt.Printf("Enter Password: ")
			fmt.Scanf("%s \n", &enter_password)

			demo_user_r := &User{enter_name, enter_password, 1}
			err_3 := Dbm.Insert(demo_user_r)
			fmt.Println(err_3)
		}
	}
	db.Close()
	// закрыть БД!!!
}
