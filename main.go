package main

import (
	"fmt"
	"sort"
	"strings"
)

func showTop(s string) {

	words := strings.Fields(s)
	cache := map[string]int{}

	for _, word := range words {
		if val, ok := cache[word]; ok {
			cache[word] = val + 1
		} else {
			cache[word] = 1
		}
	}

	type tempKv struct {
		Key   string
		Value int
	}

	var rank []tempKv

	for k, v := range cache {
		rank = append(rank, tempKv{k, v})
	}

	sort.Slice(rank, func(i, j int) bool {
		return rank[i].Value > rank[j].Value
	})

	for i, kv := range rank {
		if i < 10 {
			fmt.Printf("%s, %d\n", kv.Key, kv.Value)
		}
	}
}

func main() {
	showTop("yay lol yay hehe")
}
