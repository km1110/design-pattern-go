package main

import "github.com/km1110/design-pattern-go/template_method/display"

func main() {
	d1 := display.NewCharDiplay("H")
	d2 := display.NewStringDisplay("Hello, world.")

	display.NewDisplay(d1)
	display.NewDisplay(d2)
}
