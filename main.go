package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Book struct. structs serve as models in go

type Books struct {
	Id string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
  FirstName string `json:"firstname"`
  LastName string `json:"lastname"`
}

var books []Books 

func main(){

	//initialie router 
	router := mux.NewRouter()

	//mock data
	books = append(books, Books{
		Id:"1", 
		Isbn: "243232", 
		Title: "Book One",
		Author: &Author{FirstName: "Jon", LastName: "Akinde"}})
	books = append(books, Books{
		Id:"2", 
		Isbn: "254332", 
		Title: "Book Tow",
		Author: &Author{FirstName: "Jed", LastName: "Akinde"}})
	
	books = append(books, Books{
		Id:"3", 
		Isbn: "34567", 
		Title: "Book Three",
		Author: &Author{FirstName: "Janet", LastName: "Akinde"}})
	
	books = append(books, Books{
		Id:"4", 
		Isbn: "56789", 
		Title: "Book Four",
		Author: &Author{FirstName: "Joel", LastName: "Akinde"}})

	//router handler  
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

//router handler function have to take req qnd res like in js
func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _,v  :=range books{
		if(v.Id == params["id"]){
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	json.NewEncoder(w).Encode(&Books{})

}
func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var book Books
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.Id = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)


}
func updateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params :=mux.Vars(r)
	for k, v :=range books {
		if v.Id == params["id"] {
		books = append(books[:k], books[k+1:]... )
		var book Books
		_ = json.NewDecoder(r.Body).Decode(&book)
		book.Id = params["id"]
		books = append(books, book)
		json.NewEncoder(w).Encode(book)
		return
	}
	}
}
func deleteBook(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type","application/json")
params :=mux.Vars(r)
for k, v :=range books {
	if v.Id == params["id"] {
	books = append(books[:k], books[k+1:]... )
	break
}
	json.NewEncoder(w).Encode(books)
	// getBooks(w, r)
}
}
