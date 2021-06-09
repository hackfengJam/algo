package segmenttree

type Merger interface {
	Merge(a E, b E) E
}
type E interface {
}
