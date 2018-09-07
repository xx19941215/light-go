package main

import (
	"sort"
	"fmt"
	"math/rand"
)

type Hero struct {
	Name string
	Age int
}

//声明一个 Hero 结构体切片类型
type HeroSlice []Hero

func main() {
	var intSlice = []int{0, -1, 7, 10, 90}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄|%d", rand.Intn(10)),
			Age: rand.Intn(10),
		}

		heroes = append(heroes, hero)
	}

	for _, v := range(heroes) {
		fmt.Println(v)
	}

	sort.Sort(heroes)

	fmt.Println("排序后")
	for _, v := range(heroes) {
		fmt.Println(v)
	}
}

type Student struct {
	Name string
	Age int
	Score float64
}

func (hs HeroSlice) Len() int {
	return len(hs)
}

func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}
