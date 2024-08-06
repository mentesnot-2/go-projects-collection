package services

import (
	"github.com/mentesnot-2/library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book) error
	RemoveBook(ID int) error
	BorrowedBooks(bookId,memberId int) error
	ReturnBook(bookId,memberId int) error
	ListAvailableBooks() ([]models.Book, error)
	ListBorrowedBooks() (memberId int)

}


type Library struct {
	Books map[int]models.Book
	Members map[int]models.Member
}

func (l *Library) AddBook(book models.Book) error {
	l.Books[book.ID] = book
	return nil
}

func (l *Library) RemoveBook(ID int) error {
	delete(l.Books, ID)
	return nil
}
func (l *Library) BorrowedBooks(bookId,memberId int) error {
	book := l.Books[bookId]
	member := l.Members[memberId]
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	return nil
}

func (l *Library) ReturnBook(bookId,memberId int) error {
	book := l.Books[bookId]
	member := l.Members[memberId]
	for i, b := range member.BorrowedBooks {
		if b.ID == book.ID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break

		}
	}
	return nil
}

func (l *Library) ListAvailableBooks() ([]models.Book, error) {
	books:=[]models.Book{}
	for _,book:=range l.Books {
		if book.Status == "Available" {
			books = append(books,book)
		}
	}
	return books,nil
}

func (l *Library) ListBorrowedBooks() ([]models.Book,error) {
	books:=[]models.Book{}

	for _,book:=range l.Books {
		if book.Status == "Borrowed" {
			books = append(books,book)
		}
	}
	return books,nil
}