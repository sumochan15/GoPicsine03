package main
import (
        "fmt"
        "piscine"
)
func main() {
        fmt.Println(piscine.IsAlpha("Hello! How are you?"))
        fmt.Println(piscine.IsAlpha("HelloHowareyou"))
        fmt.Println(piscine.IsAlpha("What's this 4?"))
        fmt.Println(piscine.IsAlpha("What is this あ pen?"))
        fmt.Println(piscine.IsAlpha("This is a pen"))
        fmt.Println(piscine.IsAlpha("Whatsthis4"))
        fmt.Println(piscine.IsAlpha(""))
}