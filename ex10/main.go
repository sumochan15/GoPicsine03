package main
import (
        "fmt"
        "piscine"
)
func main() {
        fmt.Println(piscine.IsUpper("HELLO"))
        fmt.Println(piscine.IsUpper("HELLO!"))
        fmt.Println(piscine.IsUpper("あいうえお!"))
        fmt.Println(piscine.IsUpper("hello"))
        fmt.Println(piscine.IsUpper("HEllo"))
        fmt.Println(piscine.IsUpper("HEllあ"))
        fmt.Println(piscine.IsUpper(""))
}