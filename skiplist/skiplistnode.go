package skiplist

type skipListLevel struct {
	forward *SkipListNode
	span    int
}

type SkipListNode struct {
	data     int
	level    []skipListLevel
	backward *SkipListNode
}
