package skiplist_test

import (
	"fmt"
	"skiplist/skiplist"
	"testing"
)

func TestRandomLevel(t *testing.T) {
	cnt := make(map[int]int)
	for i := 0; i < 10000; i++ {
		level := skiplist.RandomLevel()
		if _, ok := cnt[level]; !ok {
			cnt[level] = 0
		} else {
			cnt[level] += 1
		}
	}

	fmt.Println("cnt:", cnt)
}

type compareInt int

func (c compareInt) SkipListNodeCompare(data interface{}) int {
	if n, ok := data.(compareInt); ok {
		return int(c - n)
	}
	panic("unexpected type")
}

func TestSkipList_Insert(t *testing.T) {
	l := skiplist.CreateSkipList()

	for i := 123; i >= 0; i-- {
		l.Insert(compareInt(i))
	}

	l.Each(func(v interface{}) {
		fmt.Print(v, ",")
	})
	fmt.Println()

	l.Reach(func(v interface{}) {
		fmt.Print(v, ",")
	})

}
