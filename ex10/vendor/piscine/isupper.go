package piscine

func isupper(s rune) bool {
		return ('A' <= s && s <= 'Z')
}

func strlen(s string) (len int) {
		for range s {
				len++
		}
		return
}

func IsUpper(s string) bool {

		r_s := []rune(s)
		for i, len := 0, strlen(s); i < len; i++ {
				if !isupper(r_s[i]) {
						return false
				}
		}
		return true
}