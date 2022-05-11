package piscine

func strlen(s string) (len int) {
		for range s {
				len++
		}
		return
}

func LastRune(s string) rune {
		len := strlen(s)
		if len > 0 {
				return []rune(s)[len-1]
		}
		return 0
}