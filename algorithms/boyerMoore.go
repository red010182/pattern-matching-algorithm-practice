package algorithms

import (
	"fmt"
	"time"
)

func printCharMap(charMap map[byte][]int) {
	fmt.Println("Bad character map:")
	for char := range charMap {
		fmt.Printf("%c, %v\n", char, charMap[char])
	}
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func BoyerMoore(T string, P string) (int, int) {
	if len(T) == 0 || len(P) == 0 {
		return -1, 0
	}

	lenP := len(P)

	// start := time.Now()
	// ┌──────────────────────────────────────────────────────────────────────────────┐
	// │ 	build bad character map                                                  		│
	// └──────────────────────────────────────────────────────────────────────────────┘
	badCharMap := make(map[byte][]int)
	for i := 0; i < lenP; i++ {
		char := P[i]
		charIndexArray := badCharMap[char]
		badCharMap[char] = append(charIndexArray, i)
	}
	// fmt.Printf("=> bad char map use %v\n", time.Since(start))
	// ┌──────────────────────────────────────────────────────────────────────────────┐
	// │ 	calcualte goodSuffixMaxMove                                                 │
	// └──────────────────────────────────────────────────────────────────────────────┘
	// start = time.Now()
	goodSuffixMaxMove := 0
	k := lenP - 1
	for ; k > 0; k-- {
		// if P[0:k] == P[lenP-k:lenP] {
		// 	break
		// }
		if P[0:k] == P[lenP-k:lenP] {
			break
		}
	}
	goodSuffixMaxMove = lenP - k
	// fmt.Printf("=> K use %v\n", time.Since(start))

	// ┌──────────────────────────────────────────────────────────────────────────────┐
	// │ 	calculate good suffix                                                    		│
	// └──────────────────────────────────────────────────────────────────────────────┘
	// start = time.Now()
	var goodSuffixTable []int = make([]int, lenP)
	for i := lenP - 2; i >= 0; i-- { // goodSuffixTable at i = lenP - 1 must be 0, therefore skip
		suffix := P[i+1:]
		// fmt.Println("target suffix", suffix, ", i: ", i)

		j := 0
		to := i + 1
		from := i + 1 - len(suffix)
		found := false
		for ; j < lenP-2*(len(suffix))+1; j++ {
			str := P[from-j : to-j]
			// fmt.Println("target str", str, ", from ", from, " to ", to, "j: ", j)
			if suffix == str { // suffix = prefix
				goodSuffixTable[i] = i - (from - j) + 1
				found = true
				break
			}
		}
		if !found {
			goodSuffixTable[i] = goodSuffixMaxMove
		}
	}
	// fmt.Printf("=> good suffix table use %v\n", time.Since(start))

	// printCharMap(badCharMap)
	// fmt.Println("Good suffix table:", goodSuffixTable)

	// ┌──────────────────────────────────────────────────────────────────────────────┐
	// │ 	BM search                                                          					│
	// └──────────────────────────────────────────────────────────────────────────────┘
	time.Now()
	compareTimes := 0
	i := 0
	for i < len(T)-lenP+1 {
		j := lenP - 1
		for j >= 0 {
			compareTimes++
			char := T[i+j]
			// fmt.Printf("compare char: %c, at index: %d\n", char, i+j)
			if char == P[j] {
				j--
			} else {
				charIndexArray, charExist := badCharMap[char]
				var badCharJump int
				if !charExist {
					badCharJump = lenP
				} else {
					for k := len(charIndexArray) - 1; k >= 0; k-- {
						distance := j - charIndexArray[k]
						if distance > 0 {
							badCharJump = distance
							break
						}
					}
				}
				i += max(badCharJump, goodSuffixTable[j])
				break
			}
		}
		if j == -1 {
			// fmt.Printf("found pattern at %d, compare %d times\n", i, compareTimes)
			// fmt.Printf("=> BM search use %v\n", time.Since(start))
			return i, compareTimes
		}
	}
	// fmt.Printf("=> BM search use %v\n", time.Since(start))
	return -1, compareTimes
}
