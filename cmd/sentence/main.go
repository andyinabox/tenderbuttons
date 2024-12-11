package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/andyinaobox/tenderbuttons/pkg/chains"
)

const pathToCorpus = "corpus/tb.txt"

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Usage: sentence <min words> <max words>")
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

	file, err := os.Open(pathToCorpus)
	if err != nil {
		panic(err)
	}

	chain := chains.NewChain(2)
	chain.Build(file)

	sentence := chain.Generate(minWords, maxWords)

	fmt.Println(sentence)
}
