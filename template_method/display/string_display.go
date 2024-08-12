package display

import "fmt"

type stringDisplay struct {
	s string
	w int
}

func NewStringDisplay(s string) IDisplay {
	w := len(s)
	return &stringDisplay{s: s, w: w}
}

func (sd stringDisplay) Open() {
	printLine(sd.w)
}

func (sd stringDisplay) Print() {
	fmt.Println("|" + sd.s + "|")
}

func (sd stringDisplay) Close() {
	printLine(sd.w)
}

func printLine(w int) {
	fmt.Print("+")
	for i := 0; i < w; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}
