package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"maps"
	"slices"

	"github.com/samber/lo"
)

/*
generate using csvkit

wget https://standards-oui.ieee.org/oui/oui.csv
echo "p,n" > mac.txt
csvcut -c "Assignment,Organization Name" oui.csv | tail -n +2 >> mac.txt
csvjson mac.txt > mac.json
rm mac.txt
*/

//go:embed mac.json
var data []byte

var Prefix = map[string]string{}

func init() {
	var pn []struct {
		P, N string
	}
	if err := json.Unmarshal(data, &pn); err != nil {
		panic(err)
	}
	for _, v := range pn {
		Prefix[v.P] = v.N
	}
}

func main() {
	fmt.Println(len(Prefix))
	fmt.Println(len(lo.Uniq(slices.Collect(maps.Values(Prefix)))))
}
