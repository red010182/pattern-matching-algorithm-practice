package algorithms

func getHash(str string) uint {
	var hash uint = 0
	N := len(str)
	for i := 0; i < N; i++ {
		hash += uint(str[i]) << (N - i - 1)
	}
	return hash
}

func RobinKarp(T string, P string) (int, int) {
	if len(T) == 0 || len(P) == 0 {
		return -1, 0
	}

	patternHash := getHash(P)
	N := len(P)
	compareTimes := 0

	var text string
	var textHash uint
	for i := 0; i < len(T)-N+1; i++ {
		compareTimes++
		text = T[i : N+i]
		if i == 0 {
			textHash = getHash(text)
		} else {
			rearHash := uint(T[N+i-1])
			frontHash := uint(T[i-1]) << (N - 1)
			textHash = (textHash-frontHash)<<1 + rearHash
		}
		if textHash == patternHash {
			// fmt.Printf("found pattern at %d, compare %d times\n", i, compareTimes)
			if text == P {
				return i, compareTimes
			}
		}
	}
	// fmt.Println("not found")
	return -1, compareTimes
}
