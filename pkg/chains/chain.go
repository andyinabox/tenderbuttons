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

const maxLoops = 10

// Chain contains a map ("chain") of prefixes to a list of suffixes.
// A prefix is a string of prefixLen words joined with spaces.
// A suffix is a single word. A prefix can have multiple suffixes.
type Chain struct {
	prefixLen int
	// main markov chain. note that the key is actuall a string representation of a prefix
	chain map[string][]string
	// contains prefixes that begin with a capital latter (stored as strings)
	startPrefixes []string
	// the key here is a single token, used to lookup possible prefixes that begin
	// with the given token
	prefixLookup map[string][]string
}

// NewChain returns a new Chain with prefixes of prefixLen words.
func NewChain(prefixLen int) *Chain {
	return &Chain{prefixLen, make(map[string][]string), make([]string, 0), make(map[string][]string)}
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

		// add to main chain
		c.chain[key] = append(c.chain[key], s)

		// add to start prefixes if first char is capital letter
		if unicode.IsUpper([]rune(key)[0]) {
			c.startPrefixes = append(c.startPrefixes, key)
		}

		// add key to a lookup which will allow looking up a prefix
		// key using the first token in the prefix
		first := p.First()
		c.prefixLookup[first] = append(c.prefixLookup[first], key)

		p.Shift(s)
	}
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

func (c *Chain) generate(n int, startPrefix Prefix, startWords []string) string {
	log.Debug("generate chain", "startPrefix", startPrefix, "startWords", startWords)
	var words []string
	var p Prefix

	outerLoops := 0

	// outer loop will keep trying to make a sentence less than n words
outer:
	for {

		outerLoops++

		// reset to start conditions
		words = make([]string, len(startWords))
		copy(words, startWords)
		p = make(Prefix, len(startPrefix))
		copy(p, startPrefix)

		// escape hatch to avoid infinite loops
		if outerLoops > maxLoops {
			log.Warn("reached max outer loops for generating chain, resetting", "pfx", startPrefix, "startWords", startWords)
			p = c.getStartPrefix()
			startWords = []string{}
		}

		log.Debug("start sentence builing loop", "pfx", p, "words", words)

		// build sentence
		for i := 0; i < n; i++ {
			// get a sice of possible next tokens
			choices := c.chain[p.String()]

			// bail out if there are no choices
			if len(choices) == 0 {
				break
			}

			// get a random token
			next := choices[rand.Intn(len(choices))]

			// append that random token to our list of words
			words = append(words, next)

			// finish if the last word has a period
			if strings.Contains(next, ".") {
				break outer
			}

			// shift token to the end of our prefix before the next loop
			p.Shift(next)
		}

		log.Debug("reached the end of outer loop without sendtence completion. retrying")
	}

	return strings.Join(words, " ")
}

func (c *Chain) getStartPrefix() Prefix {
	p1 := make(Prefix, c.prefixLen)
	p2 := NewPrefix(c.startPrefixes[rand.Intn(len(c.startPrefixes))])
	copy(p1, p2)
	return p1
}

func (c *Chain) getTokenPrefix(tok string) (p1 Prefix, ok bool) {

	prefixes, ok := c.prefixLookup[tok]

	if !ok || len(prefixes) < 1 {
		return p1, false
	}

	// copy prefix
	p1 = make(Prefix, c.prefixLen)
	p2 := NewPrefix(prefixes[rand.Intn(len(prefixes))])
	copy(p1, p2)
	ok = true

	return
}
