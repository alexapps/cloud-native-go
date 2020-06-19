package storage

import (
	model "github.com/alexapps/cloud-native-go/model"
)

// Book -
type Book interface {
	Get() (*model.Book, error)
	Create(*model.Book) (bool, string, error)
	Delete(string) (bool, error)
}
