package trie

import (
	"fmt"
)

const (
	endOfWord = rune(0x7fffffff)
)

// Trie implements a dead simple trie data structure.
type Trie struct {
	data []rune
	next []*Trie
}

// New returns a new instance of a `Trie` data structure.
func New() *Trie {
	return &Trie{
		data: []rune{},
		next: []*Trie{},
	}
}

func newNode(w []rune) *Trie {
	return &Trie{
		data: w,
		next: []*Trie{},
	}
}

// Add adds single slice of runes to the current Trie data structure.
func (t *Trie) add(w []rune) error {
	// If data[...] is empty - add the node and we are done.
	if len(t.data) == 0 && len(t.next) == 0 {
		t.data = w
		return nil
	}

	// If data[] starts with any amount of w[], forward w and add it to a new
	// *Trie node (one may exist - revisit this).
	for i, r := range w {
		// Detect no more data, that is: w is common with data and there are no
		// chars of data left to consume -- create a new trie or merge into
		// one of the Tries that exist.
		if i >= len(t.data) {
			for _, dt := range t.next {
				if dt.Merge(w[i:]) == nil {
					// Done - merge worked!
					return nil
				}
			}

			// Merge didnt work - add it to this Trie.
			t.next = append(t.next, newNode(w[i:]))
			return nil
		}

		// Iteration case - we match nodes so far and there is data left.
		if r == t.data[i] {
			continue
		}

		// Mismatch with letters left in w - split here. The parent will have
		// two tries attached here - one with the data up til now + the contents
		// of w[i:] and another with this node further qualified.
		if r == endOfWord && t.data[i] == endOfWord {
			// No op -- word already exists
			return nil
		} else if r == endOfWord {
			break
		} else {
			wNode := newNode(w[i:])
			oNode := newNode(t.data[i:])
			oNode.next = t.next
			t.data = t.data[:i]
			t.next = []*Trie{oNode, wNode}

		}
		return nil
	}

	// w (abc) was a complete subset of t.data (abcdef) - we need to split at
	// this point and separate the words.
	oldNode := newNode(t.data[len(w)-1:])
	oldNode.next = t.next

	endNode := newNode([]rune{endOfWord})

	t.data = w[:len(w)-1]
	t.next = []*Trie{endNode, oldNode}
	return nil
}

func (t *Trie) Add(w []rune) error {
	return t.add(append(w, endOfWord))
}

// Merge merges the source trie into the current one, returns error if it is not
// a legal merge.
func (t *Trie) Merge(w []rune) error {
	if t.data[0] != w[0] {
		return fmt.Errorf("merge failed")
	}
	return t.add(w)
}

func (t *Trie) allWords(prefix []rune) ([][]rune, error) {
	me := append(prefix, t.data...)
	ret := [][]rune{}
	if len(me) > 0 && me[len(me)-1] == endOfWord {
		ret = append(ret, me[:len(me)-1])
	}
	for _, st := range t.next {
		words, err := st.allWords(me)
		if err != nil {
			return nil, err
		}
		ret = append(ret, words...)
	}
	return ret, nil
}

func (t *Trie) AllWords() ([][]rune, error) {
	return t.allWords([]rune{})
}

// func (t *Trie) dump(l string, p []rune) {
// 	log.Printf("%s%s%s\n", l, string(p), string(t.data))
// 	for _, st := range t.next {
// 		st.dump(l+".", append(p, t.data...))
// 	}
// }
