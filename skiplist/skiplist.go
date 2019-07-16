package skiplist

import (
	"math/rand"
	"time"
)

const (
	MaxLevel = 32
	LevelPR  = 0.25
)

type SkipList struct {
	head   *SkipListNode
	tail   *SkipListNode
	length int
	level  int
}

func CreateSkipList() *SkipList {
	headLevel := make([]skipListLevel, MaxLevel)
	head := SkipListNode{
		nil,
		headLevel,
		nil,
	}

	tailLevel := make([]skipListLevel, MaxLevel)
	tail := SkipListNode{
		nil,
		tailLevel,
		&head,
	}

	l := &SkipList{
		&head,
		&tail,
		0,
		1,
	}
	return l
}

//插入前提是没有相同的元素存在
func (list *SkipList) Insert(data Compare) *SkipListNode {
	var (
		update [MaxLevel]*SkipListNode
	)
	//查询需要更新的节点
	x := list.head
	for i := list.level - 1; i >= 0; i-- {
		for {
			if x.level[i].forward == nil || x.level[i].forward.data.SkipListNodeCompare(data) > 0 {
				break
			} else {
				x = x.level[i].forward
			}
		}
		update[i] = x
	}

	//比原来高层设置为head
	level := RandomLevel()
	if level > list.level {
		for i := list.level; i < level; i++ {
			update[i] = list.head
		}
		list.level = level
	}

	arr := make([]skipListLevel, level)
	node := &SkipListNode{
		data,
		arr,
		nil,
	}
	//插入节点
	for i := 0; i < level; i++ {
		node.level[i].forward = update[i].level[i].forward
		update[i].level[i].forward = node
	}

	if update[0] == list.head {
		node.backward = nil
	} else {
		node.backward = update[0]
	}

	if node.level[0].forward != nil {
		node.level[0].forward.backward = node
	} else {
		list.tail = node
	}

	return node
}

//正向遍历
func (list *SkipList) Each(fn func(v interface{})) {
	cur := list.head.level[0].forward
	for {
		if cur == nil {
			break
		}
		fn(cur.data)
		cur = cur.level[0].forward
	}
}

//反向遍历
func (list *SkipList) Reach(fn func(v interface{})) {
	cur := list.tail
	for {
		if cur == nil {
			break
		}
		fn(cur.data)
		cur = cur.backward
	}
}

//返回层数， 概率为p^level
func RandomLevel() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	level := 1
	for {
		if r.Float32() < LevelPR && level < MaxLevel {
			break
		}
		level = level + 1
	}
	return level
}
