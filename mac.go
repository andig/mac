package mac

import (
	_ "embed"
	"encoding/json"
	"maps"
	"sync"
	"unique"
)

/*
generate using csvkit

wget https://standards-oui.ieee.org/oui/oui.csv
echo "p,n" > mac.txt
csvcut -c "Assignment,Organization Name" oui.csv | tail -n +2 >> mac.txt
csvjson mac.txt > mac.json
rm mac.txt
*/

var (
	//go:embed mac.json
	data []byte

	mu     sync.Mutex
	prefix map[string]unique.Handle[string]
)

func load() {
	mu.Lock()
	defer mu.Unlock()

	prefix = make(map[string]unique.Handle[string])
	var pn []struct {
		P, N string
	}

	if err := json.Unmarshal(data, &pn); err != nil {
		panic(err)
	}

	for _, v := range pn {
		prefix[v.P] = unique.Make(v.N)
	}
}

func Prefixes() map[string]unique.Handle[string] {
	if prefix == nil {
		load()
	}

	return maps.Clone(prefix)
}

func Vendor(key string) string {
	if prefix == nil {
		load()
	}

	if h, ok := prefix[key]; ok {
		return h.Value()
	}

	return ""
}
