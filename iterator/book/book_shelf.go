package book

type IBookShelf interface {
	GetBookAt(i int) Book
	AppendBook(b Book)
	GetLength() int
}

type bookShelf struct {
	books []Book
	last  int
}

func NewBookShelf(l int) IBookShelf {
	return &bookShelf{books: make([]Book, 0, l), last: 0}
}

func (bs *bookShelf) GetBookAt(i int) Book {
	return bs.books[i]
}

func (bs *bookShelf) AppendBook(b Book) {
	bs.books = append(bs.books, b)
	bs.last++
}

func (bs *bookShelf) GetLength() int {
	return bs.last
}
