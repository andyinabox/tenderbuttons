package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/andyinaobox/tenderbuttons/pkg/chains"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Usage: sentence <max words> <path/to/corpus.txt>")
		return
	}

	maxWords, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}

	path := args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	chain := chains.NewChain(2)
	chain.Build(file)

	sentence := chain.Generate(maxWords)

	fmt.Println(sentence)
}
