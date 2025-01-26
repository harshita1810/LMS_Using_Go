package library

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Library holds a collection of books
type Library struct {
	Books []Book
}

// NewLibrary creates a new Library instance
func NewLibrary() *Library {
	return &Library{}
}

// AddBook adds a new book to the library
func (l *Library) AddBook(book Book) {
	l.Books = append(l.Books, book)
	fmt.Println("Book added successfully!")
}

// RemoveBook removes a book from the library by ISBN
func (l *Library) RemoveBook(isbn string) {
	for i, book := range l.Books {
		if book.ISBN == isbn {
			l.Books = append(l.Books[:i], l.Books[i+1:]...)
			fmt.Println("Book removed successfully!")
			return
		}
	}
	fmt.Println("Book not found!")
}

// SearchBook searches for a book by title
func (l *Library) SearchBook(title string) {
	for _, book := range l.Books {
		if strings.EqualFold(book.Title, title) {
			fmt.Printf("Book found: %s\n", book)
			return
		}
	}
	fmt.Println("Book not found!")
}

// ListBooks lists all the books in the library
func (l *Library) ListBooks() {
	if len(l.Books) == 0 {
		fmt.Println("No books in the library.")
		return
	}
	fmt.Println("Books in the library:")
	for _, book := range l.Books {
		fmt.Println(book)
	}
}

// Menu provides a menu-driven interface for the Library Management System
func (l *Library) Menu() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nLibrary Menu:")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Search Book")
		fmt.Println("4. List Books")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("\nAdding a new book...")
			fmt.Print("Title: ")
			scanner.Scan()
			title := scanner.Text()
			fmt.Print("Author Name: ")
			scanner.Scan()
			author := scanner.Text()
			fmt.Print("ISBN: ")
			scanner.Scan()
			isbn := scanner.Text()
			fmt.Print("Available or Not (true/false): ")
			var available bool
			fmt.Scan(&available)

			book := Book{Title: title, Author: author, ISBN: isbn, Available: available}
			l.AddBook(book)
		case 2:
			fmt.Println("\nRemoving a book...")
			fmt.Print("Enter ISBN to remove: ")
			scanner.Scan()
			isbn := scanner.Text()
			l.RemoveBook(isbn)
		case 3:
			fmt.Println("\nSearching for a book...")
			fmt.Print("Enter Title to search: ")
			scanner.Scan()
			title := scanner.Text()
			l.SearchBook(title)
		case 4:
			fmt.Println("\nListing all books...")
			l.ListBooks()
		case 5:
			fmt.Println("Exiting the program. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
