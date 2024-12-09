package chains

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"unicode"
)

// Chain contains a map ("chain") of prefixes to a list of suffixes.
// A prefix is a string of prefixLen words joined with spaces.
// A suffix is a single word. A prefix can have multiple suffixes.
type Chain struct {
	chain         map[string][]string
	startPrefixes []Prefix
	prefixLen     int
}

// NewChain returns a new Chain with prefixes of prefixLen words.
func NewChain(prefixLen int) *Chain {
	return &Chain{make(map[string][]string), make([]Prefix, 0), prefixLen}
}

// Build reads text from the provided Reader and
// parses it into prefixes and suffixes that are stored in Chain.
func (c *Chain) Build(r io.Reader) {
	br := bufio.NewReader(r)
	p := make(Prefix, c.prefixLen)
	for {
		var s string
		if _, err := fmt.Fscan(br, &s); err != nil {
			break
		}
		key := p.String()
		c.chain[key] = append(c.chain[key], s)

		p.Shift(s)
	}

	// create collection of starting prefixes
	for prefix := range c.chain {
		// true if first letter is capitalized
		if unicode.IsUpper([]rune(prefix)[0]) {
			c.startPrefixes = append(c.startPrefixes, NewPrefix(prefix))
		}
	}
}

func (c *Chain) BuildFromString(s string) {
	c.Build(strings.NewReader(s))
}

// Generate returns a string of at most n words generated from Chain.
func (c *Chain) Generate(n int) string {
	p := c.getStartPrefix()
	words := []string(p)

	for i := 0; i < n; i++ {
		choices := c.chain[p.String()]
		if len(choices) == 0 {
			break
		}
		next := choices[rand.Intn(len(choices))]
		words = append(words, next)
		p.Shift(next)

		// finish if the last word has a period
		if strings.Contains(next, ".") {
			break
		}
	}
	return strings.Join(words, " ")
}

func (c *Chain) getStartPrefix() Prefix {
	return c.startPrefixes[rand.Intn(len(c.startPrefixes))]
}
