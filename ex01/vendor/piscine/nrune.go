package piscine

// func strlen(s string) (len int) {
// 		for range s {
// 				len++
// 		}
// 		return
// }

// func NRune(s string, n int) rune {
// 		if s == "" || strlen(s) < n || n <= 0 {
// 				return 0
// 		}
// 		return []rune(s)[n-1]
// }

func NRune(s string, n int) rune {
	for i, r := range s{
		if i+1 == n {
			return r
		}
	}
	return 0
}