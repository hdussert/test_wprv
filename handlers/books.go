package handlers

import (
	"net/http"
	"io/ioutil"
	DB "../db"
	"encoding/json"
	"fmt"
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
	dateStart := params["date_start"]
	dateEnd := params["date_end"]
	search := params["search"]

	if dateStart != "" {

	}

	if dateEnd != "" {

	}

	if search != "" {

	}
	data := []byte{}
	query := "SELECT * FROM books ORDER BY id"
	db := DB.OpenConnection()
	err := db.QueryRow(`SELECT COALESCE (array_to_json(array_agg(row_to_json(res))), '[]') FROM (` + query + `) AS res;`).Scan(&data)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

	defer db.Close()
}

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
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}

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
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}

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
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}

func FilterBooks(w http.ResponseWriter, r *http.Request, params map[string]string) {
	
	keys := r.URL.Query()
	dateStart := keys.Get("date_start")
	dateEnd := keys.Get("date_end")
	search := keys.Get("search")

	if dateStart == "" {
		dateStart = "0001-01-01"
	} 
	if dateEnd == "" {
		dateEnd = "2200-01-01"
	}

	searchString := "%"+search+"%"
	query := "SELECT * FROM books WHERE (date >= $1 AND date <= $2 AND (title LIKE $3 OR description LIKE $3)) ORDER BY id"

	data := []byte{}
	db := DB.OpenConnection()
	err := db.QueryRow(`SELECT COALESCE (array_to_json(array_agg(row_to_json(res))), '[]') FROM (` + query + `) AS res;`, dateStart, dateEnd, searchString).Scan(&data)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

	defer db.Close()
}