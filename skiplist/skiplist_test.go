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

func TestSkipList_Insert(t *testing.T) {
	l := skiplist.CreateSkipList()
	l.Insert(12)
	l.Insert(13)
	l.Insert(14)
	l.Insert(1)
	l.Insert(2)

}
