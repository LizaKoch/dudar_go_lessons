package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name      string
	age       uint16  // целове положительное число
	money     int16   // целове число
	avgGrade  float32 // число с плавающей точкой
	isStudent bool    // булево значение
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money equal: %d", u.name, u.age, u.money)
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, false}
	fmt.Fprintf(w, bob.getAllInfo())
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is page about author")
}

func contactsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is contact's page")
}

func HandleRequests() {
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/about/", aboutPage)
	http.HandleFunc("/contacts/", contactsPage)

	http.ListenAndServe(":8080", nil)
}

func main() {
	//var bob User = User{"Bob", 25, -50, 4.2, false}
	//bob = User{name: "Alex"}                // с ключами
	HandleRequests()
}
