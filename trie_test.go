package trie

////////////////////////////////////////////////////////////////////////////////

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

////////////////////////////////////////////////////////////////////////////////
func makeRuneLists(rs ...interface{}) [][]rune {
	ret := [][]rune{}
	for _, v := range rs {
		r, ok := v.([]rune)
		if ok {
			ret = append(ret, r)
		}
	}
	return ret
}
func TestTrieBasic(t *testing.T) {
	for _, tc := range []struct {
		words [][]rune
	}{
		{words: makeRuneLists("a", "b", "c", "ab")},
		{words: makeRuneLists("hi_", "hel")},
		{words: makeRuneLists("hi_", "hi_th")},
		{words: makeRuneLists("hi_", "hi_th", "hi_there")},
		{words: makeRuneLists("hi_", "hi_there", "hi_th")},
		{words: makeRuneLists("abc", "def", "ab")},
		{words: makeRuneLists("012345", "01234")},
		{words: makeRuneLists("012345", "0123567", "01299")},
	} {
		trie := New()
		for _, w := range tc.words {
			trie.Add(w)
		}

		words, err := trie.AllWords()
		assert.Equal(t, err, nil)
		assert.Equal(t, len(words), len(tc.words))
		for _, w := range tc.words {
			assert.Contains(t, words, w)
			log.Printf("Found %s\n", string(w))
		}
	}
}
