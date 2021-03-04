package graph

import (
    "errors"
    "fmt"
)

// 稀疏图（邻接表）
type SparseGraph struct {
    n, m     int  // n 为顶点 V; m 为边 E
    directed bool // 是否为有向图
    g        [][]*Edge
}

func NewSparseGraph(n int, directed bool) (g *SparseGraph) {
    g = &SparseGraph{
        n:        n,
        m:        0,
        directed: directed,
        g:        make([][]*Edge, 0, n),
    }
    for i := 0; i < n; i++ {
        g.g = append(g.g, make([]*Edge, 0))
    }
    return
}

func (g *SparseGraph) V() int {
    return g.n
}

func (g *SparseGraph) E() int {
    return g.m
}

func (g *SparseGraph) validateVertex(v int) error {
    if v < 0 || v >= g.V() {
        return errors.New(fmt.Sprintf("vertex %d is invalid", v))
    }
    return nil
}

func (g *SparseGraph) AddEdge(v, w int, weight float64) {
    // TODO g.validateVertex(v)
    // TODO g.validateVertex(w)

    if g.HasEdge(v, w) {
        return
    }

    g.g[v] = append(g.g[v], &Edge{
        a:      v,
        b:      w,
        weight: weight,
    })

    // v != w 排除自环
    if v != w && !g.directed {
        g.g[w] = append(g.g[w], &Edge{
            a:      w,
            b:      v,
            weight: weight,
        })
    }

    g.m++
}

func (g *SparseGraph) HasEdge(v, w int) bool {
    // TODO g.validateVertex(v)
    // TODO g.validateVertex(w)

    for i := 0; i < len(g.g[v]); i++ {
        if g.g[v][i].Other(v) == w {
            return true
        }
    }
    return false
}

func (g *SparseGraph) Adj(v int) []*Edge {
    // TODO g.validateVertex(v)
    return g.g[v]
}
