package tests

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{"A": 21, "C": 3, "B": 46}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}
