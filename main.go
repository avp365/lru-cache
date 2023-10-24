package main

import (
	"container/list"
	"fmt"
)

type CacheData struct {
	Key   string
	Value string
}

func main() {
	cap := 5
	list := list.New()
	e := list.PushFront(CacheData{"key", "value1"})

	fmt.Println("len\n", list.Len())
	fmt.Println("e\n", e)

}
