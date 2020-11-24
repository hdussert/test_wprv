package handlers

import (
	"net/http"
	"fmt"
)

// DeleteBooks : Takes an array of ids and delete the corresponding books
func DeleteBooks(w http.ResponseWriter, r *http.Request, params map[string]string) {
	
	var ids []int
	ParseBody(w, r, &ids)
	
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
	var queryParams []interface{}
	queryParams = append(queryParams, res)

	PostPutDeleteHandler(w, r, query, queryParams)
}