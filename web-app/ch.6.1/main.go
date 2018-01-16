// Exercise code for Chapter 6.1
// Purpose is show how to use sessions and cookies in Go
package main

import (
	"fmt"
	"time"
	"net/http"
	"html/template"
	"strings"
)

var t *template.Template
var data interface{}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	var err error
	t, err = template.ParseFiles("index.gtpl", "cookies.gtpl")
	checkError(err)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := t.ExecuteTemplate(w, "index", data)
	checkError(err)
}

// route /setcookie handler
func setCookieHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("expires_time: ", r.Form["expires_time"][0])
	current_time := time.Now().Format("15:04:05.000000 -0700 UTC")
	fmt.Println("time now: ", current_time)
	expires := strings.Join(r.Form["expires_date"], "") + " " + current_time
	fmt.Println("expires: ", expires)
	expires_time, _ := time.Parse("2006-01-02 15:04:05.000000 -0700 UTC", expires)
	fmt.Println("time now: ", expires_time)

	cookie := http.Cookie{
		Name: strings.Join(r.Form["name"], ""),
		Value: strings.Join(r.Form["value"], ""),
		Expires: expires_time,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/showcookie", http.StatusSeeOther)
}

func fetchCookieHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Name string
		Value string
		Expires time.Time
	}
	for _, cookie := range r.Cookies() {
		data.Name = cookie.Name
		data.Value = cookie.Value
		data.Expires = cookie.Expires
		fmt.Println(cookie.Expires)
	}
	t.ExecuteTemplate(w, "cookies", data)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/setcookie", setCookieHandler)
	http.HandleFunc("/showcookie", fetchCookieHandler)
	err := http.ListenAndServe(":9090", nil)
	checkError(err)
	fmt.Println("It works well.")
}
