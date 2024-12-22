package util

func InvertListValuedMap[K, V comparable](m map[K][]V) map[V][]K {
	inverted := make(map[V][]K)

	for key, l := range m {
		for _, value := range l {
			if _, found := inverted[value]; !found {
				inverted[value] = make([]K, 0)
			}

			inverted[value] = append(inverted[value], key)
		}
	}

	return inverted
}
