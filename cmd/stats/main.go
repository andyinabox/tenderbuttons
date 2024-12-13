package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"

	"github.com/charmbracelet/log"
	"gitlab.com/andyinaobox/tenderbuttons/pkg/chains"
)

const pathToCorpus = "corpus/tb.txt"

type Stats struct {
	WordCounts      []int       `json:"-"`
	CharCounts      []int       `json:"-"`
	WordCountsTable map[int]int `json:"-"`
	CharCountsTable map[int]int `json:"-"`
	AverageWords    int
	MinWords        int
	MaxWords        int
	AverageChars    int
	MinChars        int
	MaxChars        int
}

func (s *Stats) Add(sentance string) {
	wordCount := len(strings.Split(sentance, " "))
	charCount := len(sentance)

	s.WordCounts = append(s.WordCounts, wordCount)
	s.CharCounts = append(s.CharCounts, charCount)

	s.WordCountsTable[wordCount]++
	s.CharCountsTable[charCount]++

	var totalWords float64

	for _, wCount := range s.WordCounts {
		s.MinWords = int(math.Min(float64(s.MinWords), float64(wCount)))
		s.MaxWords = int(math.Max(float64(s.MaxWords), float64(wCount)))

		totalWords += float64(wCount)
	}

	s.AverageWords = int(totalWords / float64(len(s.WordCounts)))

	var totalChars float64
	for _, cCount := range s.CharCounts {
		s.MinChars = int(math.Min(float64(s.MinChars), float64(cCount)))
		s.MaxChars = int(math.Max(float64(s.MaxChars), float64(cCount)))

		totalChars += float64(cCount)
	}

	s.AverageChars = int(totalChars / float64(len(s.CharCounts)))
}

func (s *Stats) Print() {

	wctKeys := make([]int, len(s.WordCountsTable))
	for k := range s.WordCountsTable {
		wctKeys = append(wctKeys, k)
	}
	sort.Ints(wctKeys)

	for _, k := range wctKeys {
		if k == 0 {
			continue
		}

		c := ""
		for i := 0; i < s.WordCountsTable[k]; i++ {
			c = c + "+"
		}
		fmt.Printf("%d:%s\n", k, c)

	}

	b, err := json.MarshalIndent(*s, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func NewStats() *Stats {
	return &Stats{
		WordCountsTable: make(map[int]int),
		CharCountsTable: make(map[int]int),
		MinWords:        math.MaxInt,
		MinChars:        math.MaxInt,
	}
}

var min, max, prefixLen, iterations int
var verbose bool

func init() {

	flag.IntVar(&min, "min", 3, "minimum word count")
	flag.IntVar(&max, "max", 35, "maximum word count")
	flag.IntVar(&prefixLen, "p", 2, "prefix length")
	flag.IntVar(&iterations, "i", 1000, "iterations to run for stats")
	flag.BoolVar(&verbose, "v", false, "run in verbose mode")

	flag.Parse()

	if verbose {
		log.SetLevel(log.DebugLevel)
	}

}

func main() {

	log.Info("initiating stats", "min", min, "max", max, "prefixLen", prefixLen, "iterations", iterations)

	file, err := os.Open(pathToCorpus)
	if err != nil {
		panic(err)
	}

	chain := chains.NewChain(&chains.Config{
		PrefixLength: prefixLen,
		MinWords:     min,
		MaxWords:     max,
	})

	log.Info("building chain")
	chain.Build(file)
	stats := NewStats()

	log.Info("generating sentences")
	for i := 0; i < iterations; i++ {
		sentence := chain.Generate()
		fmt.Println(sentence)
		stats.Add(sentence)
	}

	stats.Print()

}
