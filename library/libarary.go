package library

import (
	"fmt"

	"github.com/Standartenfuhrer/simple-library/domain"
)

type Library struct {
	Books   []*domain.Book
	Readers []*domain.Reader

	lastBookID   int
	lastReaderID int
}

func (l *Library) AddReader(firstName, lastName string) *Reader {
	l.lastReaderID++
	newReader := &domain.Reader{
		ID:        l.lastReaderID,
		FirstName: firstName,
		LastName:  lastName,
		IsActive:  true,
	}
	l.Readers = append(l.Readers, newReader)
	fmt.Printf("Зарегистрирован новый читатель: %s\n", newReader)
	return newReader
}

func (l *Library) AddBook(year int, title, author string) *domain.Book {
	l.lastBookID++

	newBook := &domain.Book{
		ID:       l.lastBookID,
		Year:     year,
		Title:    title,
		Author:   author,
		IsIssued: false,
	}
	l.Books = append(l.Books, newBook)
	fmt.Printf("Добавлена новая книга: %s\n", newBook)
	return newBook
}

func (l *Library) FindBookById(id int) (*domain.Book, error) {
	flag := false
	for i := 0; i < len(l.Books); i++ {
		if i == id-1 {
			flag = true
		}
	}
	if flag {
		return l.Books[id-1], nil
	}
	return nil, fmt.Errorf("книга с ID %d не найдена", id)
}

func (l *Library) FindReaderById(id int) (*domain.Reader, error) {
	flag := false
	for i := 0; i < len(l.Readers); i++ {
		if i == id-1 {
			flag = true
		}
	}
	if flag {
		return l.Readers[id-1], nil
	}
	return nil, fmt.Errorf("читатель с ID %d не найден", id)
}

func (l *Library) IssueBookToReader(bookId, readerId int) error {
	book, err := l.FindBookById(bookId)
	if book == nil {
		return err
	}
	reader, err := l.FindReaderById(readerId)
	if reader == nil {
		return err
	}
	book.IssueBook(reader)
	return nil
}
