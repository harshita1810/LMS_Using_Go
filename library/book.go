package library

import "fmt"

// Book represents a book in the library
type Book struct {
	Title     string
	Author    string
	ISBN      string
	Available bool
}

// String formats the book's details
func (b Book) String() string {
	availability := "Not Available"
	if b.Available {
		availability = "Available"
	}
	return fmt.Sprintf("Title: %s, Author: %s, ISBN: %s, Availability: %s", b.Title, b.Author, b.ISBN, availability)
}
