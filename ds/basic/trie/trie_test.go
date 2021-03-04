package trie

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestT(t *testing.T) {
	trie := NewTrie()
	trie.AddWord("bad")
	trie.AddWord("dad")
	trie.AddWord("mad")

	ast := assert.New(t)

	ctx := context.Background()
	ast.Equal(trie.Search(ctx, "pad"), false)
	ast.Equal(trie.Search(ctx, "bad"), true)
	ast.Equal(trie.Search(ctx, ".ad"), true)
	ast.Equal(trie.Search(ctx, "b.."), true)
	ast.Equal(trie.Search(ctx, "."), false)

}
