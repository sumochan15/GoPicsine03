package piscine

// func countelems(elems []string) (count int) {
// 		for range elems {
// 				count++
// 		}
// 		return
// }

// func Join(strs []string, sep string) string {
// 		var str3 string = ""
// 		count := countelems(strs)
// 		for i := 0; i < count-1; i++ {
// 				str3 += strs[i]
// 				str3 += sep
// 		}

// 		if count >= 1 {
// 				str3 += strs[count-1]
// 		}

// 		return str3
// }

func Join(strs []string, sep string) string {
	var str_box string

	for i, str := range strs{
		if i != 0{
			str_box += sep
		}
		str_box += str
	}
	return str_box
}