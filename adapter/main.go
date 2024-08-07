package main

import "github.com/km1110/design-pattern-go/adapter/banner"

func main() {
	b := banner.NewBanner("Hello")
	pb := banner.NewPrintBanner(b)

	pb.PrintWeak()
	pb.PrintStrong()

}
