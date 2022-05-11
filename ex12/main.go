package main
import (
        "fmt"
        "piscine"
)
func main() { 
	fmt.Println(piscine.IsPrintable("Hello")) 
	fmt.Println(piscine.IsPrintable("Aello123"))
	fmt.Println(piscine.IsPrintable("12345abc"))
	fmt.Println(piscine.IsPrintable("12345aBc"))
	fmt.Println(piscine.IsPrintable(""))
	fmt.Println(piscine.IsPrintable("あいうえお"))
	fmt.Println(piscine.IsPrintable("あello"))
	fmt.Println(piscine.IsPrintable("Aello\n"))
	fmt.Println(piscine.IsPrintable("Hello\n"))
}