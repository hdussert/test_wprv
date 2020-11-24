package handlers

// Book : Book struct
type Book struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Autor string `json:"autor"`
	Date string `json:"date"`
}