package handlers

import (
	"net/http"
	"io/ioutil"
	DB "../db"
	"encoding/json"
)

// UpdateBook : Update an existant book
func UpdateBook(w http.ResponseWriter, r *http.Request, params map[string]string) {
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
	
	query := `UPDATE books SET description=$1, title=$2, autor=$3, date=$4 WHERE id=$5`
	
	db := DB.OpenConnection()
	_, err = db.Exec(query, book.Description, book.Title, book.Autor, book.Date, book.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}