package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/andyinaobox/tenderbuttons/pkg/chains"
)

const pathToCorpus = "corpus/tb.txt"

type Stats struct {
	WordCounts   []int `json:"-"`
	CharCounts   []int `json:"-"`
	AverageWords int
	MinWords     int
	MaxWords     int
	AverageChars int
	MinChars     int
	MaxChars     int
}

func (s *Stats) Add(sentance string) {
	s.WordCounts = append(s.WordCounts, len(strings.Split(sentance, " ")))
	s.CharCounts = append(s.CharCounts, len(sentance))

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
	b, err := json.MarshalIndent(*s, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func NewStats() *Stats {
	return &Stats{
		MinWords: math.MaxInt,
		MinChars: math.MaxInt,
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		fmt.Println("Usage: stats <min words> <max words> <iterations>")
		return
	}

	minWords, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}

	maxWords, err := strconv.Atoi(args[1])
	if err != nil {
		panic(err)
	}

	iterations, err := strconv.Atoi(args[2])
	if err != nil {
		panic(err)
	}

	file, err := os.Open(pathToCorpus)
	if err != nil {
		panic(err)
	}

	chain := chains.NewChain(2)
	chain.Build(file)
	stats := NewStats()

	for i := 0; i < iterations; i++ {
		sentence := chain.Generate(minWords, maxWords)
		fmt.Println(sentence)
		stats.Add(sentence)
	}

	stats.Print()

}
