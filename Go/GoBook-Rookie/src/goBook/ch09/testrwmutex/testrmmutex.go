package main

import (
	"fmt"
	"sync"
)

var (
	count      int
	countGurad sync.RWMutex
)

func getVal() int {
	countGurad.RLock()

	defer countGurad.RUnlock()

	return count
}

func main() {
	fmt.Println(getVal())
}
