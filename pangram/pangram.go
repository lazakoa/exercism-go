package pangram

import "strings"

func IsPangram(input string) bool {
	set := NewAlphabetSet()
	upper := strings.ToUpper(input)
	for _, char := range upper {
		flag, ok := set[char]
		if ok {
			if flag == false {
				set[char] = true
			}
		}

	}
	count := 0
	for _, v := range set {
		if v {
			count += 1
		}
	}
	if count == 26 {
		return true
	}
	return false
}

type AlphabetSet map[rune]bool

func NewAlphabetSet() AlphabetSet {
	return AlphabetSet{
		'A': false,
		'B': false,
		'C': false,
		'D': false,
		'E': false,
		'F': false,
		'G': false,
		'H': false,
		'I': false,
		'J': false,
		'K': false,
		'L': false,
		'M': false,
		'N': false,
		'O': false,
		'P': false,
		'Q': false,
		'R': false,
		'S': false,
		'T': false,
		'U': false,
		'V': false,
		'W': false,
		'X': false,
		'Y': false,
		'Z': false,
	}
}
