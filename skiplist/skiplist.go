package skiplist

import (
	"fmt"
	"math/rand"
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
		0,
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

//查找是否存在
func (list *SkipList) Exist(data Compare) bool {
	curLevel := list.level - 1
	cur := list.head
	for {
		if cur.level[curLevel].forward != nil &&
			cur.level[curLevel].forward.data.SkipListNodeCompare(data) == 0 {
			return true
		}

		if cur.level[curLevel].forward != nil &&
			cur.level[curLevel].forward.data.SkipListNodeCompare(data) < 0 {
			cur = cur.level[curLevel].forward
		} else {
			curLevel--
			if curLevel < 0 {
				return false
			}
		}
	}
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

func (list *SkipList) DebugOut() {
	for i := 0; i < list.level; i++ {
		fmt.Println("level:", i)
		cur := list.head.level[i].forward
		for {
			if cur == nil {
				break
			}
			fmt.Print(cur.data, ",")
			cur = cur.level[i].forward
		}
		fmt.Println()
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
	level := 1
	for {
		p := rand.Float32()
		if p < LevelPR && level < MaxLevel {
			level = level + 1
		} else {
			break
		}
	}
	return level
}
