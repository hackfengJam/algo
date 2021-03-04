package trie

import "context"

const wildcard = '.'

type Trie interface {
	Search(ctx context.Context, text string) bool
	AddWord(txt string)
}

type BasicTrie struct {
	Root  *TrieNode
	Count int
}

func NewTrie() Trie {
	return NewBasicTrie()
}

func NewBasicTrie() *BasicTrie {
	return &BasicTrie{
		Root: &TrieNode{
			Index: 0,
			Nexts: NewTrieRoute(),
		},
	}
}

// 不能并发调用
func (a *BasicTrie) AddWord(text string) {
	txt := []rune(text)

	if len(txt) == 0 {
		return
	}

	var (
		entry *TrieNode
		ok    bool
		level int
	)
	cur := a.Root
	for i, w := range txt {
		level = i + 1
		if entry, ok = cur.Nexts[w]; !ok {
			a.Count++
			entry = &TrieNode{
				Index:       a.Count,
				ParentIndex: cur.Index,
				Word:        w,
				Level:       level,
				Nexts:       NewTrieRoute(),
			}
			cur.Nexts[w] = entry
		}
		cur = entry
	}
	cur.End = true
}

func (a *BasicTrie) Search(ctx context.Context, text string) bool {
	txt := []rune(text)
	return a.match(ctx, a.Root, txt, 0)
}

func (a *BasicTrie) match(ctx context.Context, node *TrieNode, txt []rune, index int) bool {
	if index == len(txt) {
		return node.End
	}

	c := txt[index]
	var (
		entry *TrieNode
		ok    bool
	)
	if c != wildcard {
		if entry, ok = node.Nexts[c]; !ok {
			return false
		}
		return a.match(ctx, entry, txt, index+1)
	}
	for _, entry = range node.Nexts {
		if a.match(ctx, entry, txt, index+1) {
			return true
		}
	}
	return false
}
