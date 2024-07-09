//go:build !solution

package speller




var int_to_str map[int64]string = map[int64]string{
	0:"zero",
	1:"one",
	2:"two",
	3:"three",
	4:"four",
	5:"five",
	6:"six",
	7:"seven",
	8:"eight",
	9:"nine",
	10:"ten",
	11:"eleven",
	12:"twelve",
	13:"thirteen",
	14:"fourteen",
	15:"fifteen",
	16:"sixteen",
	17:"seventeen",
	18:"eighteen",
	19:"nineteen",
	20:"twenty",
	30:"thirty",
	40:"forty",
	50:"fifty",
	60:"sixty",
	70:"seventy",
	80:"eighty",
	90:"ninety",
	100:"hundred",
	1000:"thousand",
	1_000_000:"million",
	1_000_000_000:"billion",
}

func spell_sup(n int64, order int64) string {
	var str string


	if del := n/100; del > 0 {
		str+=int_to_str[del] + " hundred"
		n = n%100
		if n != 0 {
			str += " "
		}
	}

	if n != 0 {
		if v, ok := int_to_str[n]; ok {
			str += v
		} else {
			str += int_to_str[n/10*10]
			n = n % 10
			if n != 0 {
				str += "-"
				str += int_to_str[n]
			}
		}
	}
	if order != 1 {
		str += " " + int_to_str[order]
	}
	return str

}

func Spell(n int64) string {
	var str string
	var del int64 = 1_000_000_000

	if n < 0 {
		str += "minus "
		n*=-1
	}

	if v, ok := int_to_str[n]; ok {
		if n < 100 {
			return str + v
		} else {
			return str + "one " + v
		}
	}

	if n < 999 {
		return str + spell_sup(n, 1)
	}

	for n / del ==0 {
		del /= 1000
	}


	for n != 0 {

		str += spell_sup(n/del, del) + " "

		n = n % del
		del = del / 1000

	}

	return str[:len(str)-1]
}
