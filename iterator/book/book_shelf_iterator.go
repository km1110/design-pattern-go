package book

import "fmt"

type IBookShelfIterator interface {
	HasNext() bool
	Next() Book
}

type bookShelfIterator struct {
	bs    IBookShelf
	index int
}

func NewBookShelfIterator(bs IBookShelf) IBookShelfIterator {
	return &bookShelfIterator{bs: bs, index: 0}
}

func (bsi *bookShelfIterator) HasNext() bool {
	if bsi.index < bsi.bs.GetLength() {
		return true
	} else {
		return false
	}
}

func (bsi *bookShelfIterator) Next() Book {
	if !bsi.HasNext() {
		fmt.Println("Error")
	}

	book := bsi.bs.GetBookAt(bsi.index)
	bsi.index++
	return book
}
