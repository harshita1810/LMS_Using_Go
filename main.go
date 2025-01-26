package main

import (
	"library-management/library"
)

func main() {
	libraryManager := library.NewLibrary()
	libraryManager.Menu()
}
