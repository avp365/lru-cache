package main

import (
	"container/list"
	"fmt"
)

type CacheData struct {
	Key  string
	Data string
}

func main() {
	//cap := 5
	list := list.New()
	list.PushFront(CacheData{"key1", "value1"})
	f := list.Front()

	fmt.Println("data\n", f.Value.(CacheData))

}
