package main

import (
	"fmt"
	"sort"
)

type Sort []string

func (s Sort) Len() int {
	return len(s)
}

func (s Sort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Sort) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	test := []string{"AA", "a", "aa", "1", "222", "333", "444"}
	sort.Sort(Sort(test))
	fmt.Println(test)
}
