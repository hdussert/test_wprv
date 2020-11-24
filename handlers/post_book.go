package handlers

import (
	"net/http"
	"io/ioutil"
	DB "../db"
	"encoding/json"
)

// PostBook : Add a new book to the librairie
func PostBook(w http.ResponseWriter, r *http.Request, params map[string]string) {
	
	rByte, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var book Book
	err = json.Unmarshal(rByte, &book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	query := `INSERT INTO books (title, description, autor, date) VALUES ($1, $2, $3, $4)`
	
	db := DB.OpenConnection()
	_, err = db.Exec(query, book.Title, book.Description, book.Autor, book.Date)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}