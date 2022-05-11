package piscine

func strlen(s string) (len int) {
		for range s {
				len++
		}
		return
}

func ToLower(s string) string {
		r_s := []rune(s)

		for i, s_len := 0, strlen(s); i < s_len; i++ {
				if 'A' <= r_s[i] && r_s[i] <= 'Z' {
						r_s[i] += 32
				}
		}
		return string(r_s)
}