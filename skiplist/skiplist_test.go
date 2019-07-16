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
		if c-n > 0 {
			return 1
		}
		if c-n < 0 {
			return -1
		}

		return 0
	}
	panic("unexpected type")
}

func TestSkipList_Insert(t *testing.T) {
	l := skiplist.CreateSkipList()

	for i := 1; i <= 100000; i++ {
		l.Insert(compareInt(i))
	}

	//l.Each(func(v interface{}) {
	//	fmt.Print(v, ",")
	//})
	//fmt.Println()
	//
	//l.Reach(func(v interface{}) {
	//	fmt.Print(v, ",")
	//})
	//fmt.Println()

	//l.DebugOut()
	for i := 100000; i > 0; i-- {
		if !l.Exist(compareInt(i)) {
			t.Error("exist is no pass")
		}
	}

}
