package domain

import "fmt"

type Book struct {
	ID       int
	Year     int
	Title    string
	Author   string
	IsIssued bool
	ReaderId *int
}

func (b Book) String() string {
	status := ""
	if b.IsIssued {
		status = "используется."
		return fmt.Sprintf("ID: %d, %s (%s %d), книга %s читателем %d", b.ID, b.Title, b.Author, b.Year, status, *b.ReaderId)
	} else {
		status = "не используется."
		return fmt.Sprintf("ID: %d, %s (%s %d), книга %s", b.ID, b.Title, b.Author, b.Year, status)
	}

}

func (b *Book) IssueBook(r *Reader) error {
	if b.IsIssued {
		return fmt.Errorf("книга %s уже на руках у читателя с ID %d", b.Title, b.ReaderId)
	}
	b.IsIssued = true
	b.ReaderId = &r.ID

	return nil
}

func (b *Book) ReturnBook() error {
	if !b.IsIssued {
		return fmt.Errorf("книга %s уже в библиотеке", b.Title)
	}
	b.IsIssued = false
	b.ReaderId = nil

	return nil
}

type Reader struct {
	ID        int
	FirstName string
	LastName  string
	IsActive  bool
}

func (r *Reader) Deactivate() {
	r.IsActive = false
}

func (r *Reader) Activate() {
	r.IsActive = true
}

func (r Reader) String() string {
	status := ""
	if r.IsActive {
		status = "активен."
	} else {
		status = "неактивен."
	}
	return fmt.Sprintf("Пользователь %s %s, ID: %d, пользователь %s", r.FirstName, r.LastName, r.ID, status)
}
