package main

import (
	"fmt"
	"skiplist/linkedlist"
)
func main() {
	var l  linkedlist.Linklist
	l.InsertHead("123")
	l.InsertHead("456")
	l.Next(func(v interface{}){
		fmt.Println(v)
	})

}
