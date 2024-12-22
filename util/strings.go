package util

func FindSubStrings(str string, substr string) []int {
	i, j := 0, 0
	found := make([]int, 0)

	for i < len(str) {
		if str[i] == substr[j] {
			i++
			j++

			if j == len(substr) {
				found = append(found, i - j)
				i -= j - 1
				j = 0
			}
		} else {
			if j > 0 {
				i -= j - 1
				j = 0
			} else {
				i++
			}
		}
	}

	return found
}

func ReverseString(str string) string {
	reversed := ""

	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}

	return reversed
}
