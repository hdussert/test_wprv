package handlers

import (
	"net/http"
	"io/ioutil"
	DB "../db"
	"encoding/json"
	"log"
)

type Book struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Autor string `json:"autor"`
	Date string `json:"date"`
}

func PostBook(w http.ResponseWriter, r *http.Request, params map[string]string) {
	rByte, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var book Book
	err = json.Unmarshal(rByte, &book)
	log.Printf("%+v", book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	query := `INSERT INTO books (title, description, autor, date) VALUES ($1, $2, $3, $4)`
	
	db := DB.OpenConnection()
	_, err = db.Exec(query, book.Title, book.Description, book.Autor, book.Date)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}

func GetBooks(w http.ResponseWriter, r *http.Request, params map[string]string) {
	log.Printf("%+v", *r)
	query := "SELECT * FROM books"
	
	data := []byte{}

	db := DB.OpenConnection()
	log.Println("JELLOW")
	err := db.QueryRow(`SELECT COALESCE (array_to_json(array_agg(row_to_json(res))), '[]') FROM (` + query + `) AS res;`).Scan(&data)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

	defer db.Close()
}