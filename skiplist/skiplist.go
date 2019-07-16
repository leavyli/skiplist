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
