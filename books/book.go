package books

import (
	"encoding/json"
	"strings"
)

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b *Book) CategoryByLength() string {
	if b.Pages > 300 {
		return "NOVEL"
	} else {
		return "SHORT STORY"
	}
}

func NewBookFromJSON(_json string) (Book, error) {
	var book Book
	err := json.Unmarshal([]byte(_json), &book)
	return book, err
}

func (b *Book) AuthorLastName() string {
	s := strings.Split(b.Author, " ")
	return s[len(s)-1]
}
