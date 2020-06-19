package util

import (
	"encoding/json"
	"net/http"

	"github.com/alexapps/cloud-native-go/model"
)

// func (b Book) ToJSON() []byte {
// 	ToJSON, err := json.Marshal(b)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return ToJSON
// }

func FromJSON(data []byte) model.Book {
	book := model.Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

func WriteJSON(w http.ResponseWriter, i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}
