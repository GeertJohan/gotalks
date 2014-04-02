// +build OMIT

package main

import (
	"fmt"
	"github.com/ziutek/mymysql/autorc"
	_ "github.com/ziutek/mymysql/thrsafe"
)

func main() {
	// START OMIT
	conn := autorc.New("tcp", "", "localhost:3306", "user", "passwd", "database")
	stmtSelectFooBar, err := conn.Prepare(`SELECT foo, bar FROM stuff WHERE ID = ?`)
	if err != nil {
		panic(err)
	}
	rows, res, err := stmtSelectFooBar.Exec(42)
	if err != nil {
		panic(err)
	}
	for _, row := range rows {
		foo := row.Str(res.Map("foo"))
		bar := row.Uint64(res.Map("bar"))
		fmt.Println(foo, bar)
	}
	// END OMIT
}
