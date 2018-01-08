// Example code for Chapter 5.2 from "Build Web Application with Golang"
// Purpose: Use SQL driver to perform simple CRUD operations.
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_HOST     = "127.0.0.1:3306"
	DB_USER		= "root"
	DB_PASSWORD = "root"
	DB_NAME		= "laravel_test"
)

func main() {
	dbSouce := fmt.Sprintf("%v:%v@tcp(%v)/%v", DB_USER, DB_PASSWORD, DB_HOST, DB_NAME)
	db, err := sql.Open("mysql", dbSouce)
	checkErr(err)
	defer db.Close()

	fmt.Println("Inserting")
	stmt, err := db.Prepare("INSERT users SET name=?,email=?,password=?,remember_token=?,created_at=?,updated_at=?")
	checkErr(err)

	res, err := stmt.Exec("justencc", "test@email.com", "testpassword", "1", "2018-01-08", "2018-01-08")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("id of last inserted row=", id)
	fmt.Println("Updating")
	stmt, err = db.Prepare("UPDATE users SET name=? WHERE id=?")
	checkErr(err)

	res, err = stmt.Exec("justenccupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect, "row(s) changed")

	fmt.Println("Querying")
	rows, err := db.Query("SELECT * FROM users")
	checkErr(err)

	for rows.Next() {
		var id int
		var name, email, password, remember_token, created_at, updated_at string
		err = rows.Scan(&id, &name, &email, &password, &remember_token, &created_at, &updated_at)
		checkErr(err)
		fmt.Println("id | name | email | password")
		fmt.Printf("%3v | %6v | %6v | %6v\n", id, name, email, password)
	}

	fmt.Println("Deleting")
	stmt, err = db.Prepare("delete from users where id=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect, "row(s) changed")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
