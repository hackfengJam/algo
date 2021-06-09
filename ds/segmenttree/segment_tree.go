package segmenttree

import (
	"fmt"
	"strings"
)

type SegmentTree struct {
	data   []E
	tree   []E
	merger Merger
}

func NewSegmentTree(arr []E, merger Merger) (st *SegmentTree) {
	st = &SegmentTree{}

	st.merger = merger
	lnArr := len(arr)
	st.data = make([]E, lnArr)
	for i := 0; i < lnArr; i++ {
		st.data[i] = arr[i]
	}

	// 具体原理待补充
	st.tree = make([]E, 4*lnArr)

	st.buildSegmentTree(0, 0, len(st.data)-1)
	return st
}

// 在 treeIndex 的位置创建表示区间[l...r]的线段树
func (st *SegmentTree) buildSegmentTree(treeIndex, l, r int) {
	if l == r {
		st.tree[treeIndex] = st.data[l]
		return
	}
	leftTreeIndex := st.leftChild(treeIndex)
	rightTreeIndex := st.rightChild(treeIndex)

	mid := l + (r-l)/2
	st.buildSegmentTree(leftTreeIndex, l, mid)
	st.buildSegmentTree(rightTreeIndex, mid+1, r)

	st.tree[treeIndex] = st.merger.Merge(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的左孩子节点的索引
func (st *SegmentTree) leftChild(index int) int {
	return 2*index + 1
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的右孩子节点的索引
func (st *SegmentTree) rightChild(index int) int {
	return 2*index + 2
}

// 将 index 位置的值，更新为 e
func (st *SegmentTree) Set(index int, e E) {
	if index < 0 || index >= len(st.data) {
		// TODO throw error
		return
	}

	st.data[index] = e
	st.set(0, 0, len(st.data)-1, index, e)
}

// 在以 treeIndex为根的线段树中更新 index 的值为 e
func (st *SegmentTree) set(treeIndex, l, r, index int, e E) {
	if l == r {
		st.tree[treeIndex] = e
		return
	}

	leftTreeIndex := st.leftChild(treeIndex)
	rightTreeIndex := st.rightChild(treeIndex)

	mid := l + (r-l)/2
	if index >= mid+1 {
		st.set(rightTreeIndex, mid+1, r, index, e)
	} else {
		st.set(leftTreeIndex, l, mid, index, e)
	}

	// 注意父节点要更新
	st.tree[treeIndex] = st.merger.Merge(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
}

// 返回区间[queryL, queryB]的值
func (st *SegmentTree) Query(queryL, queryR int) E {
	return st.query(0, 0, len(st.data)-1, queryL, queryR)
}

// 在以 treeID 为根的线段树中[l...r]的范围里，搜索区间[queryL...queryR]的值
func (st *SegmentTree) query(treeIndex, l, r, queryL, queryR int) E {
	if l == queryL && r == queryR {
		return st.tree[treeIndex]
	}

	leftTreeIndex := st.leftChild(treeIndex)
	rightTreeIndex := st.rightChild(treeIndex)

	mid := l + (r-l)/2
	if queryL >= mid+1 {
		return st.query(rightTreeIndex, mid+1, r, queryL, queryR)
	} else if queryR <= mid {
		return st.query(leftTreeIndex, l, mid, queryL, queryR)
	}

	leftResult := st.query(leftTreeIndex, l, mid, queryL, mid)
	rightResult := st.query(rightTreeIndex, mid+1, r, mid+1, queryR)

	return st.merger.Merge(leftResult, rightResult)
}

func (st *SegmentTree) String() string {
	sb := strings.Builder{}
	sb.WriteString("[")
	ln := len(st.tree)
	for i := 0; i < ln; i++ {
		if st.tree[i] != nil {
			sb.WriteString(fmt.Sprintf("%v", st.tree[i]))
		} else {
			sb.WriteString("null")
		}
		if i != ln-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}
