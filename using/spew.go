// +build OMIT

package main

import "github.com/davecgh/go-spew/spew"

type Data struct {
	unexported    []byte
	Exported      string
	ptrToMoreData *Data
}

func main() {
	d := &Data{
		unexported: []byte{0x0, 0x42, 0x2A},
		Exported:   "foo bar",
		ptrToMoreData: &Data{
			Exported: "hi",
		},
	}
	spew.Dump(d)
}
