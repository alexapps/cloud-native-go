package storage

import (
	model "../../model"
)

// BookService -
type BookService struct {
}

// InitBookService -
func InitBookService() *BookService {
	return BookService{}
}

// MOC Books
var booksMap = map[string]Book{
	"0123456789": Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", Description: "bla bla", ISBN: "0123456789"},
	"4433444":    Book{Title: "Somthing Go", Author: "JJ Booo", Description: "Some interesting book", ISBN: "4433444"},
}

// Get -
func (b *BookService) Get(ID string) (*model.Book, error) {
	if item, ok := booksMap[ID]; ok {
		return item, nil
	}
	return nil, nil
}

// Create -
func (b *BookService) Create(newItem *model.Book) (bool, string, error) {
	if item, ok := booksMap[ID]; !ok {
		booksMap[ID] = newItem
		return true, item.ISBN, nil
	}
	return false, nil, nil
}

// Delete -
func (b *BookService) Delete(ID string) (bool, error) {
	if item, ok := booksMap[ID]; ok {
		delete(booksMap, ID)
		return true, nil
	}
	return false, nil
}
