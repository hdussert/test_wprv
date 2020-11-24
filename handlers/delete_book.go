package handlers

import (
	"net/http"
	"io/ioutil"
	DB "../db"
	"encoding/json"
	"fmt"
)

// DeleteBook : Takes an id and the corresponding book
func DeleteBook(w http.ResponseWriter, r *http.Request, params map[string]string) {
	rByte, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var id int
	err = json.Unmarshal(rByte, &id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	query := `DELETE FROM books WHERE id=$1`
	
	db := DB.OpenConnection()
	_, err = db.Exec(query, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}


// DeleteBooks : Takes an array of ids and delete the corresponding books
func DeleteBooks(w http.ResponseWriter, r *http.Request, params map[string]string) {
	rByte, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var ids []int

	err = json.Unmarshal(rByte, &ids)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	res := ""
	for i, v := range ids { // Dirty way to get {1,2,3,7} format used by ANY
		if i == 0 {
			res += `{`
		}
		res += fmt.Sprint(v)
		if i < len(ids) - 1 {
				res += ","
		}
		if i == len(ids) - 1 {
			res += `}`
		}
	}

	query := `DELETE FROM books WHERE (id IS NOT NULL AND id=ANY($1))`
	db := DB.OpenConnection()
	_, err = db.Exec(query, res)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}