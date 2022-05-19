package piscine

func StrLen(str string)(len int){
	for range str{
		len++
	}
	return
}

func CheckBase(str string)bool{
	if StrLen(str) < 2 {
		return false
	}
	
	for i1, c_word1 := range str{
		if c_word1 == '-' || c_word1 == '+'{
			return false
		}
		for _, c_word2 := range str[i1+1:]{
			if c_word1 == c_word2{
				return false
			}
		}
	}
	return true
}

func Check_S(s string,base string)bool{
	for _, v1 := range s {
		for _, v2 := range base {
			if v1 != v2 {
				continue
			}
			if v1 == v2 {
				goto check
			}
		}
		return false
		
		check:
	}
	return true
}

func ChangeAtoi(s string, base string)int{
	var atoi_s int
	var base_num int

	for _, n1 := range s {
		for i, n2 := range base{
			if n1 == n2 {
				base_num = i
			}
		}
		atoi_s = atoi_s * StrLen(base) + base_num
	}
	return atoi_s
}

func AtoiBase(s string, base string) int {
	var last_ans int
	if !Check_S(s, base) || !CheckBase(base){
		return 0
	}
	last_ans = ChangeAtoi(s, base)
	return last_ans
}