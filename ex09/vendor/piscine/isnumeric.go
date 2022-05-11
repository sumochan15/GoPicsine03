package piscine

func isdigit(s rune) bool {
		return ('0' <= s && s <= '9')
}

func strlen(s string) (len int) {
		for range s {
				len++
		}
		return
}

func IsNumeric(s string) bool {
		r_s := []rune(s)

		for i, len := 0, strlen(s); i < len; i++ {
				if !isdigit(r_s[i]) {
						return false
				}
		}
		return true
}