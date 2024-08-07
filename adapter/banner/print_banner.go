package banner

type printBanner struct {
	banner Banner
}

func NewPrintBanner(b Banner) IPrint {
	return &printBanner{b}
}

func (pb printBanner) PrintWeak() {
	pb.banner.ShowWithParen()
}

func (pb printBanner) PrintStrong() {
	pb.banner.ShowWithAster()
}
