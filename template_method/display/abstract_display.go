package display

type IDisplay interface {
	Open()
	Print()
	Close()
}

type AbstractDisplay struct {
	IDisplay
}

func NewAbstractDisplay(d IDisplay) *AbstractDisplay {
	return &AbstractDisplay{d}
}

func (ad *AbstractDisplay) Display() {
	ad.Open()
	for i := 0; i < 5; i++ {
		ad.Print()
	}
	ad.Close()
}
