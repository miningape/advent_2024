package util

func FindSubStrings(str string, substr string) []int {
	i, j := 0, 0
	found := make([]int, 0)

	for i < len(str) {
		if str[i] == substr[j] {
			i++
			j++

			if j >= len(substr) {
				found = append(found, i - len(substr))
				j = 0
			}
		} else {
			if j > 0 {
				j = 0
			} else {
				i++
			}
		}
	}
	
	return found
}