package main

import (
	"fmt"

	"github.com/km1110/design-pattern-go/iterator/book"
)

func main() {
	bs := book.NewBookShelf(3)
	bs.AppendBook(*book.NewBook("プログラミング言語Go"))
	bs.AppendBook(*book.NewBook("Java言語で学ぶデザインパターン入門"))
	bs.AppendBook(*book.NewBook("SOFT SKILL"))

	bsi := book.NewBookShelfIterator(bs)

	for bsi.HasNext() {
		b := bsi.Next()
		fmt.Println(b.GetName())
	}
}
