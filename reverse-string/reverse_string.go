package reverse

func Reverse(input string) string {
	var reversed string
	for _, char := range input {
		reversed = string(char) + reversed
	}
	return reversed
}
