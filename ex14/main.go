package main
import (
        "fmt"
        "piscine"
)
func main() {
        elems := []string{"Hello!", " How", " are", " ゆう?"}
        fmt.Println(piscine.BasicJoin(elems))
}