// +build OMIT

package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

const blob = `[
	{"Title":"Avisi", "URL":"http://avisi.nl"},
	{"Title":"Foize", "URL":"http://foize.com"}
]`

type Item struct {
	Title string
	URL   string
}

func main() {
	var items []*Item
	json.NewDecoder(strings.NewReader(blob)).Decode(&items)
	for _, item := range items {
		fmt.Printf("Title: %v URL: %v\n", item.Title, item.URL)
	}
}
