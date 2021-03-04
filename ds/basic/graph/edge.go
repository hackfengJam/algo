package graph

import "fmt"

type Edge struct {
    a      int
    b      int
    weight float64
}

func NewEdge(a, b int, weight float64) *Edge {
    return &Edge{
        a:      a,
        b:      b,
        weight: weight,
    }
}

func (e *Edge) V() int {
    return e.a
}

func (e *Edge) W() int {
    return e.b
}

func (e *Edge) Wt() float64 {
    return e.weight
}

func (e *Edge) Other(x int) int {
    // TODO assert (x==a || x == b)
    if x == e.a {
        return e.a
    }
    return e.b
}

// Comparable
func (e *Edge) CompareTo(e1 *Edge) int {
    return int(e.weight - e1.weight)
}

// toString
func (e *Edge) String() string {
    return fmt.Sprintf("%d-%d %.2f", e.a, e.b, e.weight)
}
