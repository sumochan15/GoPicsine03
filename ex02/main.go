package main
import (
		"ft"
		"piscine"
)
func main() {
		ft.PrintRune(piscine.LastRune("Hello!"))
		ft.PrintRune(piscine.LastRune("Salutあ"))
		ft.PrintRune(piscine.LastRune("Ola!"))
		ft.PrintRune(piscine.LastRune("Ola "))
		ft.PrintRune('\n')
		ft.PrintRune(piscine.LastRune(""))
}