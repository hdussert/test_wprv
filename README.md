# test_wprv

## Overview
```
.
├── database.sql
├── db
│   └── db.go
├── handlers
│   ├── book.go
│   ├── delete_book.go
│   ├── get_book.go
│   ├── post_book.go
│   └── put_book.go
├── server.go
└── setup.sh
```

## Starting

Install modules : `sh setup.sh`.  
Create the database described in `database.sql`.  
Update `db/db.go` to connect your database.  
Start the server : `go run server.go`.  

## Routes `localhost:8080/`


### GET `localhost:8080/books`
### POST `localhost:8080/books/` 
Body envoyé:
```
{
	"title": "13 Shades of Pumpkins",
	"description": "This is a description",
	"autor": "hdussert",
	"date": "2014-09-12"
}
```
### PUT `localhost:8080/books/`  
Body envoyé:
```
{
	"id": 1,
	"title": "No Shades of Pmpkins",
	"description": "This description has been changed",
	"autor": "hdussert",
	"date": "2014-09-18"
}
```
### DELETE `localhost:8080/books/`
Body envoyé:
```
[2,3]
```

### GET `localhost:8080/books/filter`
Url parameters (optional) : search, date_start, date_end  
`http://localhost:8080/books/filters?search=Pum&date_start=2005-12-01&date_end=2019-11-02`
