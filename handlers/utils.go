package handlers

import (
	"net/http"
	DB "../db"
	"io/ioutil"
	"encoding/json"
)

// GetHandler : Handles get queries
func GetHandler(w http.ResponseWriter, r *http.Request, query string) {
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

// PostPutDeleteHandler : Handles Post Put Delete
func PostPutDeleteHandler(w http.ResponseWriter, r *http.Request, query string, queryParams []interface{}) {
	db := DB.OpenConnection()
	_, err := db.Exec(query, queryParams...)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}

// GetHandler : Handles get queries
func GetFilterHandler(w http.ResponseWriter, r *http.Request, query string, queryParams []interface{}) {
	data := []byte{}
	db := DB.OpenConnection()
	err := db.QueryRow(`SELECT COALESCE (array_to_json(array_agg(row_to_json(res))), '[]') FROM (` + query + `) AS res;`, queryParams...).Scan(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

	defer db.Close()
}

// ParseBody : Yes
func ParseBody(w http.ResponseWriter, r *http.Request, s interface{}) {
	rByte, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.Unmarshal(rByte, &s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}