package display

import "fmt"

type charDisplay struct {
	ch string
}

func NewCharDiplay(ch string) IDisplay {
	return &charDisplay{ch: ch}
}

func (cd charDisplay) Open() {
	fmt.Print("<<")
}

func (cd charDisplay) Print() {
	fmt.Print(cd.ch)
}

func (cd charDisplay) Close() {
	fmt.Println(">>")
}
