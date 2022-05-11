package piscine

func check(s rune) bool {
		return (32 <= s && s <= 126)
}

func IsPrintable(s string) bool {

		for _, c := range s {
				if !check(c) {
						return false
				}
		}
		return true
}