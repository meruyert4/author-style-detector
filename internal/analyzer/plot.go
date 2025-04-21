package analyzer

import (
	"fmt"
	"sort"
)

func DisplayStemLeaf(counts []int) {
	stemMap := make(map[int][]int)

	for _, count := range counts {
		stem := count / 10
		leaf := count % 10
		stemMap[stem] = append(stemMap[stem], leaf)
	}

	stems := make([]int, 0, len(stemMap))
	for stem := range stemMap {
		stems = append(stems, stem)
	}
	sort.Ints(stems)

	for _, stem := range stems {
		leaves := stemMap[stem]
		sort.Ints(leaves)
		fmt.Printf("%2d | ", stem)
		for _, leaf := range leaves {
			fmt.Printf("%d ", leaf)
		}
		fmt.Println()
	}
}
