// Exercise code for Chapter 5.5
// Purpose is to show to use BeeDB ORM for basic CRUD operations for mysql
package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id              int `orm:"auto"`
	Name			string
	Email			string
	Password		string
	Remember_token	string
	Created_at		string
	Updated_at		string
}

func (u *User) TableName() string {
	return "users"
}

const (
	DB_HOST = "127.0.0.1:3306"
	DB_USER = "root"
	DB_PASS = "root"
	DB_NAME = "laravel_test"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getTimeStamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func init() {
	// register model
	orm.RegisterModel(new(User))

	// set default database
	dbSouce := fmt.Sprintf("%v:%v@tcp(%v)/%v", DB_USER, DB_PASS, DB_HOST, DB_NAME)
	orm.RegisterDataBase("default", "mysql", dbSouce, 30)
}

func main() {
	o := orm.NewOrm()

	user := User{
		Name: "gentcys",
		Email: "gentcys@example.com",
		Password: "password",
		Remember_token: "1",
		Created_at: getTimeStamp(),
		Updated_at: getTimeStamp(),
	}

	// insert
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// update
	user.Name = "justencc"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// read one
	u := User{Id: user.Id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	// delete
	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
