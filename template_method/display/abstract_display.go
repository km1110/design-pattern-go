package display

func NewDisplay(d IDisplay) {
	d.Open()
	for i := 0; i < 5; i++ {
		d.Print()
	}
	d.Close()
}
