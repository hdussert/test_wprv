package handlers

import (
	"net/http"
)

// UpdateBook : Update an existant book
func UpdateBook(w http.ResponseWriter, r *http.Request, params map[string]string) {

	var book Book
	ParseBody(w, r, &book)
	query := `UPDATE books SET description=$1, title=$2, autor=$3, date=$4 WHERE id=$5`

	var queryParams []interface{}
	queryParams = append(queryParams, book.Description)
	queryParams = append(queryParams, book.Title)
	queryParams = append(queryParams, book.Autor)
	queryParams = append(queryParams, book.Date)
	queryParams = append(queryParams, book.ID)
	PostPutDeleteHandler(w, r, query, queryParams)
}