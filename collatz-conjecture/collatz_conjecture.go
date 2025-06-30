package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("n is less than or equal to 0")
	}
	step := 0
	for {
		if n == 1 {
			return step, nil
		} else if n%2 == 0 {
			n /= 2
			step += 1
		} else {
			n = 3*n + 1
			step += 1
		}
	}
}
