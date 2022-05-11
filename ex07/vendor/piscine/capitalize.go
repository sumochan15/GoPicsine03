package piscine

func isupper(s rune) bool {
		return ('A' <= s && s <= 'Z')
}

func islower(s rune) bool {
		return ('a' <= s && s <= 'z')
}

func alphanumeric(s rune) bool {
		return ('a' <= s && s <= 'z' || 'A' <= s && s <= 'Z' || '0' <= s && s <= '9')
}

func strlen(s string) (len int) {
		for range s {
				len++
		}
		return
}

func Capitalize(s string) string {
		r_s := []rune(s)
		for i, len := 0, strlen(s); i < len; i++ {
				if isupper(r_s[i]) {
						r_s[i] += 32
				}
				if islower(r_s[i]) && (i == 0 || !alphanumeric(r_s[i-1])) {
						r_s[i] -= 32
				}
		}
		return string(r_s)
}