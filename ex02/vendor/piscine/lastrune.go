package piscine

// func strlen(s string) (len int) {
// 		for range s {
// 				len++
// 		}
// 		return
// }

// func LastRune(s string) rune {
// 		len := strlen(s)
// 		if len > 0 {
// 				return []rune(s)[len-1]
// 		}
// 		return 0
//

func LastRune(s string) rune {
	var r_str rune
	var count int
	for _, str := range s{
		r_str = str
		count++
	}
	if count > 0{
		return r_str
	}
	return 0
}