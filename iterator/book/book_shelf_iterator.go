package book

import (
	"errors"
)

type IBookShelfIterator interface {
	HasNext() bool
	Next() (*Book, error)
}

type bookShelfIterator struct {
	bs    IBookShelf
	index int
}

func NewBookShelfIterator(bs IBookShelf) IBookShelfIterator {
	return &bookShelfIterator{bs: bs, index: 0}
}

func (bsi *bookShelfIterator) HasNext() bool {
	return bsi.index < bsi.bs.GetLength()
}

func (bsi *bookShelfIterator) Next() (*Book, error) {
	if !bsi.HasNext() {
		return nil, errors.New("no such elemet")
	}

	book := bsi.bs.GetBookAt(bsi.index)
	bsi.index++
	return book, nil
}
