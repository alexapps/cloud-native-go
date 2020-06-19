package storage

import (
	"github.com/alexapps/cloud-native-go/model"
)

// BookService -
type BookService struct {
}

// InitBookService -
func InitBookService() *BookService {
	return &BookService{}
}

// MOC Books
var booksMap = map[string]*model.Book{
	"0123456789": &model.Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", Description: "bla bla", ISBN: "0123456789"},
	"4433444":    &model.Book{Title: "Somthing Go", Author: "JJ Booo", Description: "Some interesting book", ISBN: "4433444"},
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
	if _, ok := booksMap[newItem.ISBN]; !ok {
		booksMap[newItem.ISBN] = newItem
		return true, newItem.ISBN, nil
	}
	return false, "", nil
}

// Delete -
func (b *BookService) Delete(ID string) (bool, error) {
	if _, ok := booksMap[ID]; ok {
		delete(booksMap, ID)
		return true, nil
	}
	return false, nil
}
