package trie

type TrieRoute map[rune]*TrieNode

type TrieNode struct {
	Index       int // node的编号
	ParentIndex int
	Word        rune // 到达这个状态的 Char，其实可以不用记
	Level       int  // node的深度，eg root节点下一级节点的深度为1
	Nexts       TrieRoute
	End         bool // end 为true表示叶子节点
}

func NewTrieRoute() TrieRoute {
	return make(TrieRoute)
}
