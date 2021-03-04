package graph

type Graph interface {
    V() int
    E() int
    AddEdge(v, w int, weight float64)
    HasEdge(v, w int) bool
    Adj(v int) []*Edge
}
