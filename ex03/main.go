package main
import (
		"fmt"
		"piscine"
)
func main() {
		fmt.Println(piscine.Index("Hello!", "l"))
		fmt.Println(piscine.Index("Salut!", "alu"))
		fmt.Println(piscine.Index("Salut!", "alw"))
		fmt.Println(piscine.Index("Salut!", "!"))
		fmt.Println(piscine.Index("Ola!", "hOl"))
		fmt.Println(piscine.Index("Ola!", ""))
}
