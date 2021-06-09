package segmenttree

import (
	"fmt"
	"testing"

	"github.com/hackfengJam/gokit/itype"
)

type Sum struct {
}

func (m *Sum) Merge(a E, b E) E {
	return E(itype.Int(a) + itype.Int(b))
}

type Mer func(a E, b E) E

func TestSegmentTree(t *testing.T) {
	nums := []E{-2, 0, 3, -5, 2, -1}

	var segmentTree *SegmentTree
	segmentTree = NewSegmentTree(nums, &Sum{})

	t.Log(segmentTree.Query(0, 2)) // 1
	t.Log(segmentTree.Query(2, 5)) // -1
	t.Log(segmentTree.Query(0, 5)) // -3

	segmentTree.Set(0, E(1))
	t.Log(segmentTree.Query(0, 2)) // 4

	t.Log(fmt.Sprintf("segmentTree: %s\n", segmentTree))
}
