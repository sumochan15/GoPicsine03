package piscine

func strlen(s string) (len int) {
		for range s {
				len++
		}
		return
}

func NRune(s string, n int) rune {
		if s == "" || strlen(s) < n || n <= 0 {
				return 0
		}
		return []rune(s)[n-1]

}