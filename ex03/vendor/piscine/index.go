package piscine

func strlen(s string) (len int) {
		for range s {
				len++
		}
		return
}

func Index(s string, toFind string) int {
		rune_s := []rune(s)
		rune_t := []rune(toFind)

		if strlen(s) == 0 || strlen(toFind) == 0 || strlen(s) < strlen(toFind) {
				return -1
		}

		for s_i, s_len := 0, strlen(s); s_i < s_len; s_i++ {
		check:
				if rune_s[s_i] == rune_t[0] {
						for t_i, next_t := range rune_t {
								if rune_s[s_i+t_i] != next_t {
										s_i++
										goto check
								}
						}
						return s_i
				}
		}
		return -1
}