package piscine

func islower(s rune) bool {
		return ('a' <= s && s <= 'z')
}

func strlen(s string) (len int) {
		for range s {
				len++
		}
		return
}

func IsLower(s string) bool {
		r_s := []rune(s)

		for i, len := 0, strlen(s); i < len; i++ {
				if !islower(r_s[i]) {
						return false
				}
		}
		return true
}