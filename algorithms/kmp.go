package algorithms

func Kmp(T string, P string) (int, int) {
	if len(T) == 0 || len(P) == 0 {
		return -1, 0
	}

	var next []int = make([]int, len(P))
	//
	// next table:
	// find the length of a suffix which is also a prefix in string[0:j], for each j
	//
	preprocess := func() {
		i, j := 0, 1
		for j < len(P) {
			if P[i] != P[j] {
				next[j] = 0
				j++
			} else {
				next[j] = i + 1
				i++
				j++
			}
		}
	}
	preprocess()
	// fmt.Println("next table: ", next)

	//
	// start
	//
	i := 0
	j := 0
	compareTimes := 0
	search := false
	for i < len(T) {
		compareTimes++
		// fmt.Printf("%c, %c\n", T[i], P[j])
		if T[i] == P[j] {
			search = true
			i++
			j++
			if j == len(P) {
				// fmt.Printf("found pattern at %d, compare %d times\n", i-j, compareTimes)
				return i - j, compareTimes
			}
		} else {
			//
			// if break from a search, compare again so do not move i, otherwise, move to next
			//
			if search == false {
				i++
			}
			if j-1 < 0 {
				j = 0
			} else {
				j = next[j-1]
			}
			search = false
		}

	}
	// fmt.Println("not found")
	return -1, compareTimes
}
