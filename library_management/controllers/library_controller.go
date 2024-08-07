package controllers

import (
    "bufio"
    "fmt"
    "github.com/mentesnot-2/library_management/models"
    "github.com/mentesnot-2/library_management/services"
    "os"
    "strconv"
    // "strings"
)

func RunLibraryConsole() {
    library := services.NewLibrary()

    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Println("Library Management System")
        fmt.Println("1. Add a new book")
        fmt.Println("2. Remove an existing book")
        fmt.Println("3. Borrow a book")
        fmt.Println("4. Return a book")
        fmt.Println("5. List all available books")
        fmt.Println("6. List all borrowed books by a member")
        fmt.Println("7. Exit")
        fmt.Print("Enter your choice: ")

        scanner.Scan()
        choice, _ := strconv.Atoi(scanner.Text())

        switch choice {
        case 1:
            addBook(scanner, library)
        case 2:
            removeBook(scanner, library)
        case 3:
            borrowBook(scanner, library)
        case 4:
            returnBook(scanner, library)
        case 5:
            listAvailableBooks(library)
        case 6:
            listBorrowedBooks(scanner, library)
        case 7:
            fmt.Println("Exiting...")
            return
        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }
}

func addBook(scanner *bufio.Scanner, library services.LibraryManager) {
    fmt.Print("Enter book ID: ")
    scanner.Scan()
    id, _ := strconv.Atoi(scanner.Text())

    fmt.Print("Enter book title: ")
    scanner.Scan()
    title := scanner.Text()

    fmt.Print("Enter book author: ")
    scanner.Scan()
    author := scanner.Text()

    book := models.Book{
        ID:     id,
        Title:  title,
        Author: author,
        Status: "Available",
    }

    library.AddBook(book)
    fmt.Println("Book added successfully.")
}

func removeBook(scanner *bufio.Scanner, library services.LibraryManager) {
    fmt.Print("Enter book ID to remove: ")
    scanner.Scan()
    id, _ := strconv.Atoi(scanner.Text())

    err := library.RemoveBook(id)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Book removed successfully.")
    }
}

func borrowBook(scanner *bufio.Scanner, library services.LibraryManager) {
    fmt.Print("Enter book ID to borrow: ")
    scanner.Scan()
    bookID, _ := strconv.Atoi(scanner.Text())

    fmt.Print("Enter member ID: ")
    scanner.Scan()
    memberID, _ := strconv.Atoi(scanner.Text())

    err := library.BorrowBook(bookID, memberID)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Book borrowed successfully.")
    }
}

func returnBook(scanner *bufio.Scanner, library services.LibraryManager) {
    fmt.Print("Enter book ID to return: ")
    scanner.Scan()
    bookID, _ := strconv.Atoi(scanner.Text())

    fmt.Print("Enter member ID: ")
    scanner.Scan()
    memberID, _ := strconv.Atoi(scanner.Text())

    err := library.ReturnBook(bookID, memberID)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Book returned successfully.")
    }
}

func listAvailableBooks(library services.LibraryManager) {
    books := library.ListAvailableBooks()
    if len(books) == 0 {
        fmt.Println("No available books.")
        return
    }
    fmt.Println("Available Books:")
    for _, book := range books {
        fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
    }
}

func listBorrowedBooks(scanner *bufio.Scanner, library services.LibraryManager) {
    fmt.Print("Enter member ID: ")
    scanner.Scan()
    memberID, _ := strconv.Atoi(scanner.Text())

    books, err := library.ListBorrowedBooks(memberID)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    if len(books) == 0 {
        fmt.Println("No borrowed books for this member.")
        return
    }
    fmt.Println("Borrowed Books:")
    for _, book := range books {
        fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
    }
}
