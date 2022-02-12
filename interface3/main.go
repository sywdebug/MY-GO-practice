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

type HeroSlicec []Hero

func (hs HeroSlicec) Len() int {
	return len(hs)
}

func (hs HeroSlicec) Less(i, j int) bool {
	return hs[i].Age > hs[j].Age
}
func (hs HeroSlicec) Swap(i, j int) {
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp
}

func main() {
	var intSlice = []int{0, 4, 78, 125, 45, 12, 235, 45}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	var heroes HeroSlicec
	for i := 0; i < 10; i++ {
		var hero = Hero{
			Name: fmt.Sprintf("英雄%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}
	fmt.Println(heroes)
	for _, v := range heroes {
		fmt.Println(v)
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
	sort.Sort(heroes)
	for _, v := range heroes {
		fmt.Println(v)
	}
}
