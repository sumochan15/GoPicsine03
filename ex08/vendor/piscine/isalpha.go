package piscine

func strlen(s string) (len int) {
		for range s {
				len++
		}
		return
}

func alphanumeric(s rune) bool {
		return ('a' <= s && s <= 'z' || 'A' <= s && s <= 'Z' || '0' <= s && s <= '9')
}

func IsAlpha(s string) bool {

		r_s := []rune(s)
		for i, len := 0, strlen(s); i < len; i++ {
				if !alphanumeric(r_s[i]) {
						return false
				}
		}
		return true
}