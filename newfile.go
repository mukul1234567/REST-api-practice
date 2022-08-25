package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	ISBN   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}
type User struct {
	Srno           string `json:"srno"`
	FirstName      string `json:"userfirstname"`
	LastName       string `json:"userlastname"`
	Bookspurchased int    `json:"bookspurchased"`
	Amount         int    `json:"amount"`
}
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book
var users []User

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Date", "Wed, 16 Aug 2022 12:15:38 GMT")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// json.NewEncoder(w).Encode(&Book{})
	fmt.Fprintf(w, "No Data available")
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn((10000000)))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
	fmt.Fprintf(w, "New Entry created and added to the database")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			fmt.Fprintf(w, "Data of the current entry has been updated")
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
	fmt.Fprintf(w, "The current entry hass been deleted")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range users {
		if item.Srno == params["srno"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// json.NewEncoder(w).Encode(&Book{})
	fmt.Fprintf(w, "No Data available")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.Srno = strconv.Itoa(rand.Intn((10000000)))
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
	fmt.Fprintf(w, "New Entry created and added to the database")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range users {
		if item.Srno == params["srno"] {
			users = append(users[:index], users[index+1:]...)
			var user User
			_ = json.NewDecoder(r.Body).Decode(&user)
			users = append(users, user)
			json.NewEncoder(w).Encode(user)
			fmt.Fprintf(w, "Data of the current entry has been updated")
			return
		}
	}
	json.NewEncoder(w).Encode(users)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range users {
		if item.Srno == params["srno"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
	fmt.Fprintf(w, "The current entry hass been deleted")
}

func main() {
	r := mux.NewRouter()

	books = append(books, Book{ID: "1", ISBN: "7741852", Title: "Char",
		Author: &Author{Firstname: "Sachin", Lastname: "Shinde"}},
		Book{ID: "2", ISBN: "9963852", Title: "Pahunchaya",
			Author: &Author{Firstname: "Virat", Lastname: "Patil"}})
	users = append(users, User{Srno: "1", FirstName: "Chinmay",
		LastName: "Ponting", Bookspurchased: 0, Amount: 0},
		User{Srno: "2", FirstName: "Rushi", LastName: "Gilchrist",
			Bookspurchased: 0, Amount: 0})
	
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/users/{srno}", getUser).Methods("GET")
	r.HandleFunc("/api/users", createUser).Methods("POST")
	r.HandleFunc("/api/users/{srno}", updateUser).Methods("PUT")
	r.HandleFunc("/api/users/{srno}", deleteUser).Methods("DELETE")
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8000", r))
}
