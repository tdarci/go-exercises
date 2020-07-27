package main

import (
	"fmt"
	"log"

	"github.com/tdarci/go-exercises/books/bookdetails"
	"github.com/tdarci/go-exercises/books/booklist"
	"github.com/tdarci/go-exercises/books/simple"
)

// ***CHANGE THIS TO YOUR DIRECTORY***
const dataDirectory = "/Users/tom/src/goplain/src/github.com/tdarci/go-exercises/books/data"

func main() {
	simple.NoPipleline()
	tryBookFunctions()
}

func tryBookFunctions() {

	fmt.Println("")
	fmt.Println("Trying out book functions...")
	fmt.Println("")

	listService, err := booklist.NewService(dataDirectory)
	if err != nil {
		log.Fatalf("Unable to create list service: %s", err)
	}

	detailService := bookdetails.NewService(dataDirectory)

	authorName := "William Shakespeare"

	fmt.Println("")
	books := listService.GetByAuthor(authorName)
	fmt.Printf("Books by '%s'...\n", authorName)
	for idx, b := range books {
		fmt.Printf("* %d. Author: %s. Title: %s\n", idx, b.Author, b.Title)
	}

	if len(books) > 0 {

		firstBook, err := detailService.Get(books[0].Filename)
		if err != nil {
			log.Fatalf("Error getting book: %s", err)
		}

		defer firstBook.Close()

		start := make([]byte, 256)
		count, readErr := firstBook.Read(start)
		if readErr != nil {
			log.Fatalf("Error reading book: %s", readErr)
		}

		fmt.Printf("\nFirst %d bytes of book: %s\n", count, start)
		fmt.Println("")
	}

}
