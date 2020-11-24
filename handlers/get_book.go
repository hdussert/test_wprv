package handlers

import (
	"net/http"
	DB "../db"
)

// GetBooks : Get all books
func GetBooks(w http.ResponseWriter, r *http.Request, params map[string]string) {
	query := "SELECT * FROM books ORDER BY id"
	data := []byte{}
	db := DB.OpenConnection()
	err := db.QueryRow(`SELECT COALESCE (array_to_json(array_agg(row_to_json(res))), '[]') FROM (` + query + `) AS res;`).Scan(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

	defer db.Close()
}

// FilterBooks : Takes optional url parameters search, date_start, date_end and return the list of books matching
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
	query := "SELECT * FROM books WHERE (date >= $1 AND date <= $2 AND (title LIKE $3 OR description LIKE $3 OR autor LIKE $3)) ORDER BY id"
	var queryParams []interface{}
	queryParams = append(queryParams, dateStart)
	queryParams = append(queryParams, dateEnd)
	queryParams = append(queryParams, searchString)

	GetFilterHandler(w, r, query, queryParams)
}