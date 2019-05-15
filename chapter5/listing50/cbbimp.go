package main

import (
	"fmt"
	"sort"
)

type Ints []int

func (a Ints) Len() int           { return len(a) }
func (a Ints) Less(i, j int) bool { return a[i] < a[j] }
func (a Ints) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	arr := []int{1, 2, 3, 8, 0, 1, 2, 3}
	sort.Sort(Ints(arr))
	fmt.Println(arr)

	arrb := []int{2, 3, 4, 1, 6, 3, 99, 22, 10, 2}
	sort.Slice(arrb, func(i, j int) bool {
		return arrb[i] > arrb[j]
	})
	fmt.Println(arrb)
}
