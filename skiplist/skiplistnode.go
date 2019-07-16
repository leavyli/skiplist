package skiplist

type skipListLevel struct {
	forward *SkipListNode
	span    int
}

type Compare interface {
	SkipListNodeCompare(data interface{}) int
}

type SkipListNode struct {
	data     Compare
	level    []skipListLevel
	backward *SkipListNode
}
