package piscine

import "ft"

func StrLen(str string)(len int){
	for range str{
		len++
	}
	return
}

func NumberSign(nbr int)(signmark int){
	signmark = 1

	if nbr < 0 {
		signmark = -1
	}
	return
}

func PrintNegative(){
		ft.PrintRune('N')
		ft.PrintRune('V')
}

func CheckBase(base string)bool{
	base_len := StrLen(base)

	if base_len < 2{
		return false
	}

	for i1, c1 := range base{
		if c1 == '-' || c1 == '+'{
			return false
		}
		for _, c2 := range base[i1+1:]{
			if c1 == c2 {
				return false
			}
		}
	}
	return true
}

func ChangeBaseNumber(nbr int, base string){
	sign := NumberSign(nbr)
	base_len := StrLen(base)

	if nbr/base_len*sign == 0 {
		if nbr < 0{
			ft.PrintRune('-')
		}
		ft.PrintRune(rune(base[nbr*sign]))
		return
	}
	ChangeBaseNumber(nbr/base_len,base)
	ft.PrintRune(rune(base[nbr%base_len*sign]))
}

func PrintNbrBase(nbr int, base string) {
	if !CheckBase(base){
		PrintNegative()
		return
	}
	ChangeBaseNumber(nbr, base)
}