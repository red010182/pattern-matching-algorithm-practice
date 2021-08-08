package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"

	"pattern_match/algorithms"
	// "pattern_match/algorithms"
	"strings"
	"time"
)

const patternMinLength int = 5
const numPatterns = 10000 // get randome patterns from dict.txt

type algo func(string, string) (int, int)
type multialgo func(string, []string) (int, int, map[string]int)

func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

// get N unique patterns
func getRandomPatternFromDict(filePath string, N int) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	patternMap := make(map[string]bool)
	patterns := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.Split(text, " ")[0]
		_, exist := patternMap[text]
		if len(text) > patternMinLength && !exist {
			patternMap[text] = true
			patterns = append(patterns, text)
		}
	}

	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	np := len(patterns) // number of filtered-patterns
	for i := 0; i < N; i++ {
		index := r.Intn(np - i)
		patterns[index], patterns[np-1-i] = patterns[np-1], patterns[index]
	}
	patterns = patterns[np-N : np]

	// writeLines(patterns, "patterns.txt")
	return patterns
}

func benchmark(fn algo, name string, text string, patterns []string) {
	totalFound := 0
	totalCompareTimes := 0
	fmt.Printf("\n============= %s =============\n", name)
	start := time.Now()
	for i := 0; i < len(patterns); i++ {
		z, times := fn(text, patterns[i])
		if z != -1 {
			totalFound += 1
			totalCompareTimes += times
			// fmt.Printf("'%s' found at %d, compare %d times\n", patterns[i], z, times)
		}
	}
	fmt.Printf("\n=> %s use %v\n", name, time.Since(start))
	fmt.Printf("=> Found %d patterns, with total %d compare times\n", totalFound, totalCompareTimes)
}

func benchmarkMulti(fn multialgo, name string, text string, patterns []string) {
	fmt.Printf("\n============= %s =============\n", name)
	start := time.Now()
	totalFound, totalCompareTimes, _ := fn(text, patterns)
	fmt.Printf("\n=> %s use %v\n", name, time.Since(start))
	fmt.Printf("=> Found %d patterns, with total %d compare times\n", totalFound, totalCompareTimes)

	// foundPatterns := []string{}
	// for pattern := range foundPatternMap {
	// 	foundPatterns = append(foundPatterns, pattern)
	// }
	// writeLines(foundPatterns, name+".txt")
}

func main() {
	bytes, err := ioutil.ReadFile("./bible.txt")
	if err != nil {
		panic("read text file error")
	}
	text := string(bytes)

	patterns := getRandomPatternFromDict("./dict.txt", numPatterns)
	fmt.Printf("\nText: bible.txt (%d characters)", len(text))
	fmt.Printf("\nRandomly picked %d patterns. Each pattern has min length %d\n", len(patterns), patternMinLength)

	benchmarkMulti(algorithms.AhoCorsasick, "Aho Corasick", text, patterns)
	benchmarkMulti(algorithms.WuManber, "Wu Manber", text, patterns)

	benchmark(algorithms.BruteForce, "Brute Force", text, patterns)
	benchmark(algorithms.Kmp, "KMP", text, patterns)
	benchmark(algorithms.BoyerMoore, "Boyer Moore", text, patterns)
	benchmark(algorithms.RobinKarp, "Robin Karp", text, patterns)
	// text = "CPM_annual_conference_announce_annually"
	// patterns = []string{"announce", "annual", "annually"}

	fmt.Println("")

}
