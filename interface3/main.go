package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Name string
	Age  int
}

type HeroSlice []Hero

func (hs HeroSlice) Len() int {
	return len(hs)
}
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Name > hs[j].Name
}
func (hs HeroSlice) Swap(i, j int) {
	// temp := hs[i]
	// hs[i] = hs[j]
	// hs[j] = temp
	hs[i], hs[j] = hs[j], hs[i]
}
func main() {
	intSlice := []int{0, -1, 10, 7, 90}
	sort.Ints(intSlice)
	fmt.Println(intSlice)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	heroes := HeroSlice{}
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}
	for _, v := range heroes {
		fmt.Println("前", v)
	}
	fmt.Println()
	sort.Sort(heroes)
	for _, v := range heroes {
		fmt.Println("后", v)
	}
}
