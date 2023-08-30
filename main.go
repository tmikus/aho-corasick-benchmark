package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"text/tabwriter"
	"time"

	bs "github.com/BobuSumisu/aho-corasick"
	ak "github.com/anknown/ahocorasick"
	cf "github.com/cloudflare/ahocorasick"
	ih "github.com/iohub/ahocorasick"
	tm "github.com/tmikus/ahocorasick_rs"
)

func timeStuff(fn func()) float64 {
	start := time.Now()
	fn()
	end := time.Now()
	return float64(end.UnixNano()-start.UnixNano()) * 1e-6
}

func totalAlloc() float64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return float64(m.TotalAlloc) / (1024 * 1024 * 1024)
}

func iohub(patterns []string, input string) (float64, float64, int) {
	m := ih.NewMatcher()

	buildTime := timeStuff(func() {
		for i, pattern := range patterns {
			m.Insert([]byte(pattern), i)
		}
		m.Compile()
	})

	seq := []byte(input)
	numMatches := 0

	searchTime := timeStuff(func() {
		resp := m.Match(seq)
		for resp.HasNext() {
			items := resp.NextMatchItem(seq)
			for range items {
				numMatches++
			}
		}
		resp.Release()
	})

	return buildTime, searchTime, numMatches
}

func anknown(patterns []string, input string) (float64, float64, int) {
	runePatterns := make([][]rune, len(patterns))
	for i, pattern := range patterns {
		runePatterns[i] = []rune(pattern)
	}
	runeInput := []rune(input)

	m := new(ak.Machine)
	var matches []*ak.Term

	buildTime := timeStuff(func() {
		if err := m.Build(runePatterns); err != nil {
			panic(err)
		}
	})

	searchTime := timeStuff(func() {
		matches = m.MultiPatternSearch(runeInput, false)
	})

	return buildTime, searchTime, len(matches)
}

func bobusumisu(patterns []string, input string) (float64, float64, int) {
	var matches []*bs.Match
	var trie *bs.Trie

	buildTime := timeStuff(func() {
		trie = bs.NewTrieBuilder().AddStrings(patterns).Build()
	})

	searchTime := timeStuff(func() {
		matches = trie.MatchString(input)
	})

	return buildTime, searchTime, len(matches)
}

func cloudflare(patterns []string, input string) (float64, float64, int) {
	var m *cf.Matcher
	var matches []int
	buildTime := timeStuff(func() {
		m = cf.NewStringMatcher(patterns)
	})
	searchTime := timeStuff(func() {
		matches = m.Match([]byte(input))
	})
	return buildTime, searchTime, len(matches)
}

func tmikus(patterns []string, input string) (float64, float64, int) {
	var matches []tm.Match
	var trie *tm.AhoCorasick

	buildTime := timeStuff(func() {
		trie = tm.NewAhoCorasick(patterns)
	})

	searchTime := timeStuff(func() {
		matches = trie.FindAll(input)
	})

	return buildTime, searchTime, len(matches)
}

type testDef struct {
	name string
	fn   func([]string, string) (float64, float64, int)
}

var tests = []testDef{
	{"anknown   ", anknown},
	{"bobusumisu", bobusumisu},
	{"cloudflare", cloudflare},
	{"iohub     ", iohub},
	{"tmikus    ", tmikus},
}

func loadPatterns() []string {
	f, err := os.Open("EOWL-v1.1.2/words.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	patterns := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		patterns = append(patterns, scanner.Text())
	}
	return patterns
}

func generateSentence(patterns []string, wordCount int) string {
	r := rand.New(rand.NewSource(int64(wordCount)))
	sentence := ""
	for i := 0; i < wordCount; i++ {
		sentence += patterns[r.Intn(len(patterns))] + " "
	}
	return strings.TrimSpace(sentence)
}

func runMultipleTimes(test testDef, patterns []string, input string) (float64, float64, int, float64) {
	runTimes := 10
	buildTime := 0.0
	searchTime := 0.0
	numMatches := 0
	totalAllocations := 0.0
	for i := 0; i < runTimes; i++ {
		runtime.GC()
		allocationsBefore := totalAlloc()
		b, s, n := test.fn(patterns, input)
		buildTime += b
		searchTime += s
		numMatches += n
		totalAllocations += totalAlloc() - allocationsBefore
	}
	return buildTime / float64(runTimes), searchTime / float64(runTimes), numMatches / runTimes, totalAllocations / float64(runTimes)
}

func main() {
	patterns := loadPatterns()

	fmt.Println("Running benchmarks...")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', tabwriter.AlignRight)
	fmt.Fprintf(w, "name\tpatterns\tinput len\tbuild\tsearch\tmatches\talloc\t\n")
	for sentenceLength := 1; sentenceLength <= 65536; sentenceLength *= 2 {
		input := generateSentence(patterns, sentenceLength)
		for _, test := range tests {
			fmt.Println("Running", test.name, "with a sentence length of", sentenceLength, "words...")
			buildTime, searchTime, numMatches, totalAllocations := runMultipleTimes(test, patterns, input)
			fmt.Fprintf(w, "%s\t%d\t%d\t%.02fms\t%.02fms\t%d\t%.02fGiB\t\n", test.name, len(patterns), len(input), buildTime, searchTime, numMatches, totalAllocations)
		}
		fmt.Fprintf(w, "\t\t\t\t\t\t\t\n")
	}

	w.Flush()
}
