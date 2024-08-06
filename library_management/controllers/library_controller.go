package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/mentesnot-2/library_management/models"
	"github.com/mentesnot-2/library_management/services"
)

var book = models.Book{Status: "Available", ID: 0, Title: "The power of now", Author: "Eckhart Tolle"}
var library = services.Library{
	Books: make(map[int]models.Book),
	Members: make(map[int]models.Member),
}

func AddBook() {

	fmt.Println("Library Management System: ", library)
	if library.Books == nil {
		library.Books = make(map[int]models.Book)
	}
	var bookId = len(library.Books) + 1
	book.ID = bookId

	var title string
	var author string
	var status string

	fmt.Print("Enter the title of the book: ")
	reader:= bufio.NewReader(os.Stdin)
	title, _ = reader.ReadString('\n')
	book.Title = strings.TrimSpace(title)

	fmt.Print("Enter the author of the book: ")
	newReader:= bufio.NewReader(os.Stdin)
	author, _ = newReader.ReadString('\n')
	book.Author = author

	fmt.Print("Enter the status of the book: ")
	statusReader:= bufio.NewReader(os.Stdin)
	status, _ = statusReader.ReadString('\n')
	book.Status = status
	library.AddBook(book)
	fmt.Println("Library Management System2: ", library.Books[bookId])
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
