package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	stripped := strings.ReplaceAll(id, " ", "")
	stripped = strings.TrimSpace(stripped)
	if len(stripped) < 2 {
		return false
	}

	skip := true
	total := 0

	for i := len(stripped) - 1; i >= 0; i-- {
		letter := stripped[i]
		value, err := strconv.Atoi(string(letter))
		if err != nil {
			return false
		}

		if skip == true {
			total += value
			skip = !skip
		} else {
			value *= 2
			if value > 9 {
				value -= 9
			}
			total += value
			skip = !skip
		}
	}
	if total%10 == 0 {
		return true
	} else {
		return false
	}
}
