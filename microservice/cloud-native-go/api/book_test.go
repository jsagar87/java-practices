package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookToJSON(t *testing.T) {
	book := Book{Title: "CLoud Native GO", Author: "M.-L. Reimer", ISBN: "0123456789"}
	json := book.ToJSON()
	assert.Equal(t, `{"title":"CLoud Native GO","author":"M.-L. Reimer","isbn":"0123456789"}`,
		string(json), "Book JSON marshalling wrong")
}

func TestBookFromJSON(t *testing.T) {
	json := []byte(`{"title":"CLoud Native GO","author":"M.-L. Reimer","isbn":"0123456789"}`)
	book := FromJSON(json)
	assert.Equal(t, Book{Title: "CLoud Native GO", Author: "M.-L. Reimer", ISBN: "0123456789"}, book, "Book JSON munarshalling wrong")
}
