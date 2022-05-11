package piscine

func countelems(elems []string) (count int) {
		for range elems {
				count++
		}
		return
}

func BasicJoin(elems []string) string {
		var str_mix string = ""
		for i, count := 0, countelems(elems); i < count; i++ {

				str_mix += elems[i]
		}
		return str_mix
}