package controllers

import (
	"fmt"

	"github.com/mentesnot-2/library_management/models"
	"github.com/mentesnot-2/library_management/services"
)

func AddBook() {
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

func RemoveBook() {
	var library services.Library
	var bookId int
	fmt.Println("Enter the ID of the book you want to remove: ")
	fmt.Scanln(&bookId)

	// Check if the bookId does exist in library.Books
	if _, ok := library.Books[bookId]; !ok {
		fmt.Println("The book does not exist")
		return
	}

	library.RemoveBook(bookId)
	fmt.Println("Book removed successfully")
}

func BorrowBook() {
	var library services.Library
	var bookId int
	var memberId int
	fmt.Println("Enter the ID of the book you want to borrow: ")
	fmt.Scanln(&bookId)

	// Check if the bookId does exist in library.Books
	if _, ok := library.Books[bookId]; !ok {
		fmt.Println("The book does not exist")
		return

	}
	if library.Books[bookId].Status == "Borrowed" {
		fmt.Println("The book is already borrowed")
		return
	}

	fmt.Println("Enter the ID of the member: ")
	fmt.Scanln(&memberId)

	// Check if the memberId does exist in library.Members
	if _, ok := library.Members[memberId]; !ok {
		fmt.Println("The member does not exist")
		return
	}

	book:=library.Books[bookId]
	book.Status = "Borrowed"
	library.Books[bookId] = book

	library.BorrowBook(bookId, memberId)
	fmt.Println("Book borrowed successfully")
}

func ReturnBook() {
	var library services.Library
	var bookId int
	var memberId int

	fmt.Println("Enter the ID of the book you want to return: ")	
	fmt.Scanln(&bookId)

	// Check if the bookId does exist in library.Books
	fmt.Println("Enter the ID of the member: ")
	fmt.Scanln(&memberId)


	borrowedBooks:=library.Members[memberId].BorrowedBooks
	for _, book:=range borrowedBooks {
		if book.ID == bookId {
			book.Status = "Available"
			library.Books[bookId] = book
			library.ReturnBook(bookId, memberId)
			fmt.Println("Book returned successfully")
			return
		}
	}
}

func ListAvailableBooks() {
	var library services.Library

	books,err:= library.ListAvailableBooks()
	if err != nil {
		fmt.Println("Error listing available books")
		return
	}
	for _, book:=range books {
		fmt.Printf("Title: %s By: %s\n", book.Title, book.Author)
	}
	
}

func ListBorrowedBooks() {
	var library services.Library
	var memberId int

	fmt.Println("Enter the ID of the member: ")
	fmt.Scanln(&memberId)

	if _, ok := library.Members[memberId]; !ok {
		fmt.Println("The member does not exist")
		return 
	}
   books,err:= library.ListBorrowedBooks(memberId)
	if err != nil {
		fmt.Println("Error listing borrowed books")
	}
	for _, book:=range books {
		fmt.Printf("Title: %s By: %s\n", book.Title, book.Author)
	}
}
