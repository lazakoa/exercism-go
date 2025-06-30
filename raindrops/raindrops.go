package raindrops

import "strconv"

func Convert(number int) string {
	str := ""
	factor3 := factorN(3)
	factor5 := factorN(5)
	factor7 := factorN(7)

	if factor3(number) {
		str += "Pling"
	}
	if factor5(number) {
		str += "Plang"
	}
	if factor7(number) {
		str += "Plong"
	}

	if str == "" {
		return strconv.Itoa(number)
	} else {
		return str
	}

}

func factorN(n int) func(int) bool {
	return func(number int) bool {
		return number%n == 0
	}
}
