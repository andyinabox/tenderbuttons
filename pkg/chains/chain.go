package chains

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"unicode"

	"github.com/charmbracelet/log"
)

// Chain contains a map ("chain") of prefixes to a list of suffixes.
// A prefix is a string of prefixLen words joined with spaces.
// A suffix is a single word. A prefix can have multiple suffixes.
type Chain struct {
	chain         map[string][]string
	startPrefixes []Prefix
	prefixLookup  map[string][]Prefix // used to lookup a prefix by first term
	prefixLen     int
}

// NewChain returns a new Chain with prefixes of prefixLen words.
func NewChain(prefixLen int) *Chain {
	return &Chain{make(map[string][]string), make([]Prefix, 0), make(map[string][]Prefix), prefixLen}
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
	for prefixStr := range c.chain {
		p := NewPrefix(prefixStr)
		key := []string(p)[0]

		c.prefixLookup[key] = append(c.prefixLookup[key], p)

		// true if first letter is capitalized
		if unicode.IsUpper([]rune(prefixStr)[0]) {
			c.startPrefixes = append(c.startPrefixes, p)
		}
	}

	// for str, prefixes := range c.prefixLookup {
	// 	tokens := make([]string, len(prefixes))
	// 	for i, p := range prefixes {
	// 		tokens[i] = p.String()
	// 	}

	// 	log.Debugf("%q: %s", str, fmt.Sprintf(`"%s"`, strings.Join(tokens, `" "`)))
	// }
}

func (c *Chain) BuildFromString(s string) {
	c.Build(strings.NewReader(s))
}

// Generate returns a string of at most n words generated from Chain.
func (c *Chain) Generate(n int) string {
	p := c.getStartPrefix()
	return c.generate(n, p, []string(p))
}

func (c *Chain) GenerateFromToken(tok string, n int) string {

	p, ok := c.getTokenPrefix(tok)

	if !ok {
		p = c.getStartPrefix()
	}

	return c.generate(n, p, []string(p))
}

func (c *Chain) generate(n int, p Prefix, words []string) string {
	log.Debug("generate chain", "pfx", p, "words", words)
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
	p := make(Prefix, c.prefixLen)
	copy(p, c.startPrefixes[rand.Intn(len(c.startPrefixes))])
	return p
}

func (c *Chain) getTokenPrefix(tok string) (p Prefix, ok bool) {

	prefixes, ok := c.prefixLookup[tok]

	if !ok || len(prefixes) < 1 {
		return p, false
	}

	// copy prefix
	p = make(Prefix, c.prefixLen)
	copy(p, prefixes[rand.Intn(len(prefixes))])
	ok = true

	return
}
