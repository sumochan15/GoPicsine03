package main
import (
	"ft"
	// "github.com/42tokyo/ft"
	"piscine"
)
func main() {
		piscine.PrintNbrBase(12, "0123456789")
		ft.PrintRune('\n')
		piscine.PrintNbrBase(-9223372036854775808, "0123456789")
		ft.PrintRune('\n')
		piscine.PrintNbrBase(922337203685477587, "0123456789")
		ft.PrintRune('\n')
		piscine.PrintNbrBase(-125, "01")
		ft.PrintRune('\n')
		piscine.PrintNbrBase(125, "0")
		ft.PrintRune('\n')
		piscine.PrintNbrBase(15, "0123456789ABCDEF")
		ft.PrintRune('\n')
		piscine.PrintNbrBase(-5, "choumi")
		ft.PrintRune('\n')
		piscine.PrintNbrBase(125, "aa")
		ft.PrintRune('\n')
}