package main

import "github.com/km1110/design-pattern-go/template_method/display"

func main() {
	d1 := display.NewAbstractDisplay(display.NewCharDiplay("H"))
	d2 := display.NewAbstractDisplay(display.NewStringDisplay("Hello World."))

	d1.Display()
	d2.Display()
}
