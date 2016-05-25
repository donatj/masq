package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	s := NewScanner(r)
	p := NewParser(s)

	sc, err := p.Parse()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", sc)

	for _, tbl := range sc.Tables {
		spew.Dump(tbl)
	}

}
