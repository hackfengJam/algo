package trie_prefix

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
	ast.Equal(trie.Search(ctx, "pad-1123123"), false)
	ast.Equal(trie.Search(ctx, "bad-123123123"), true)
	ast.Equal(trie.Search(ctx, "dadbcd-123123123"), true)
	ast.Equal(trie.Search(ctx, "madabvc-123123132"), true)
	ast.Equal(trie.Search(ctx, "a-123123123"), false)

}
