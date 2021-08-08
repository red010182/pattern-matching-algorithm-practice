package algorithms

import (
	"math"
	"strings"
)

const blockSize = 2

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func findMinRightMostOccur(sp string, patterns []string) int {
	var min int = math.MaxInt64
	for n := 0; n < len(patterns); n++ {
		pattern := patterns[n]
		i := strings.Index(reverse(pattern), reverse(sp))
		if i > -1 {
			if n == 0 {
				min = i
			} else if i < min {
				min = i
			}
		}
	}
	return min
}
func WuManber(text string, patterns []string) (int, int, map[string]int) {
	shiftTable := make(map[string]int)
	lastBlock := make(map[string][]int)
	foundPatternMap := make(map[string]int)

	lmin := len(patterns[0])
	for i := 0; i < len(patterns); i++ {
		pattern := patterns[i]
		lenP := len(pattern)
		if lenP < lmin {
			lmin = lenP
		}
		for j := 0; j < lenP-blockSize+1; j++ {
			sp := pattern[j : j+blockSize]
			shiftTable[sp] = -1
		}
	}

	shiftTableMax := lmin - blockSize + 1

	for sp, _ := range shiftTable {
		shiftTable[sp] = findMinRightMostOccur(sp, patterns)
	}

	for i := 0; i < len(patterns); i++ {
		pattern := patterns[i]
		lenP := len(pattern)
		suffix := pattern[lenP-blockSize : lenP]
		lastBlock[suffix] = append(lastBlock[suffix], i)
	}

	// fmt.Println(shiftTableMax)
	// fmt.Println(shiftTable)
	// fmt.Println(lastBlock)

	i := lmin
	compareTimes := 0
	totalFound := 0
	for i < len(text)-blockSize+1 {
		compareTimes++
		sp := text[i : i+blockSize]
		jump, exist := shiftTable[sp]
		if !exist {
			jump = shiftTableMax
		}
		if jump == 0 {
			patternIndexes := lastBlock[sp]
			for j := 0; j < len(patternIndexes); j++ {
				pattern := patterns[patternIndexes[j]]
				lenP := len(pattern)
				to := i + blockSize
				from := to - lenP
				if from < 0 {
					continue
				}
				if pattern == text[from:to] {
					// fmt.Printf("found %s at %d\n", pattern, from)
					_, hasFound := foundPatternMap[pattern]
					if !hasFound {
						totalFound++
						foundPatternMap[pattern] = from
					}
				}
			}
			i += 1
		}
		i += jump
	}
	return totalFound, compareTimes, foundPatternMap
}
