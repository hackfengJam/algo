package trieprefix

import (
	"context"

	"github.com/hackfengJam/algo/ds/basic/trie"
)

type PreFixTrie struct {
	trie.BasicTrie
}

func NewTrie() trie.Trie {
	return NewPreFixTrie()
}

func NewPreFixTrie() *PreFixTrie {
	return &PreFixTrie{*trie.NewBasicTrie()}
}

func (p *PreFixTrie) Search(ctx context.Context, text string) bool {
	txt := []rune(text)
	cur := p.Root
	if len(txt) == 0 {
		return cur.End
	}
	var (
		ok    bool
		entry *trie.TrieNode
	)
	for _, v := range txt {
		if entry, ok = cur.Nexts[v]; !ok {
			break
		}
		cur = entry
	}

	return cur.End
}
