package main
import (
		"ft"
		"piscine"
)
func main() {
		ft.PrintRune(piscine.FirstRune("Hello!"))
		ft.PrintRune(piscine.FirstRune("Salut!"))
		ft.PrintRune(piscine.FirstRune("Ola!"))
		ft.PrintRune('\n')
		ft.PrintRune(piscine.FirstRune("987\n"))
		ft.PrintRune(piscine.FirstRune("\n"))
		ft.PrintRune(piscine.FirstRune("H"))
		ft.PrintRune(piscine.FirstRune(""))
		ft.PrintRune(piscine.FirstRune("あa!"))
		ft.PrintRune('\n')
}