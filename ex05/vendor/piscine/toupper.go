package piscine

func strlen(s string) (len int) {
		for range s {
				len++
		}
		return
}

func ToUpper(s string) string {
		r_s := []rune(s)

		for i, s_len := 0, strlen(s); i < s_len; i++ {
				if 'a' <= r_s[i] && r_s[i] <= 'z' {
						r_s[i] -= 32
				}
		}
		return string(r_s)
}