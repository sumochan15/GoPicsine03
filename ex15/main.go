package main
import (
        "fmt"
        "piscine"
)
func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	// toConcat := []string{"Hello!", " How\n", " are", " you?"}
	fmt.Println(piscine.Join(toConcat, ":"))
	// toConcat := []string{"あ", " い", " う", " ?"}
	// fmt.Println(piscine.Join(toConcat, ":"))
        // toConcat := []string{""}
        // fmt.Println(piscine.Join(toConcat, ":"))
}