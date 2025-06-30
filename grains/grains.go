package grains

import "errors"

func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		return 0, errors.New("number is out of range")
	}
	base := uint64(1)
	for i := 1; i < number; i++ {
		base *= uint64(2)
	}
	return base, nil
}

func Total() uint64 {
	total := uint64(0)
	for i := 1; i <= 64; i++ {
		square, _ := Square(i)
		total += square
	}
	return total
}
