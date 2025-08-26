package main

import (
	"fmt"
	"maps"
	"slices"

	"github.com/andig/mac"
	"github.com/samber/lo"
)

func main() {
	prefixes := mac.Prefixes()
	fmt.Println(len(prefixes))
	fmt.Println(len(lo.Uniq(slices.Collect(maps.Values(prefixes)))))
}
