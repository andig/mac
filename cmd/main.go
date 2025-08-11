package main

import (
	"fmt"
	"maps"
	"slices"

	"github.com/andig/mac"
	"github.com/samber/lo"
)

func main() {
	fmt.Println(len(mac.Prefix))
	fmt.Println(len(lo.Uniq(slices.Collect(maps.Values(mac.Prefix)))))
}
