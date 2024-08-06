package controllers

import (
	"fmt"

	"github.com/mentesnot-2/library_management/models"
	"github.com/mentesnot-2/library_management/services"
)

func AddBook(){
	var book models.Book
	var library services.Library
	var bookId = len(library.Books) + 1
	book.ID = bookId

	var title string
	var author string
	var status string


	fmt.Println("Enter the title of the book: ")
	fmt.Scanln(&title)	
	book.Title = title

	fmt.Println("Enter the author of the book: ")
	fmt.Scanln(&author)
	book.Author = author
	
	fmt.Println("Enter the status of the book: ")
	fmt.Scanln(&status)
	book.Status = status
	library.AddBook(book)
	fmt.Println("Book added successfully")
}