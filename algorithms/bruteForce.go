package algorithms

func BruteForce(T string, P string) (int, int) {
	if len(T) == 0 || len(P) == 0 {
		return -1, 0
	}

	i, j := 0, 0
	compareTimes := 0
	for i < len(T) {
		compareTimes++
		if j < 0 || T[i] == P[j] {
			i++
			j++
			if j == len(P) {
				// fmt.Printf("found pattern at %d, compare %d times\n", i-j, compareTimes)
				return i - j, compareTimes
			}
		} else {
			i = i - j
			j = -1
		}
	}
	// fmt.Println("not found")
	return -1, compareTimes
}
