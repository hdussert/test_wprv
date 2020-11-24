package handlers

import (
	"net/http"
)

// PostBook : Add a new book to the librairie
func PostBook(w http.ResponseWriter, r *http.Request, params map[string]string) {
	var book Book
	ParseBody(w, r, &book)
	query := `INSERT INTO books (title, description, autor, date) VALUES ($1, $2, $3, $4)`

	var queryParams []interface{}
	queryParams = append(queryParams, book.Title)
	queryParams = append(queryParams, book.Description)
	queryParams = append(queryParams, book.Autor)
	queryParams = append(queryParams, book.Date)

	PostPutDeleteHandler(w, r, query, queryParams)
}