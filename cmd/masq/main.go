package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/donatj/masq"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	s := masq.NewScanner(r)
	p := masq.NewParser(s)

	sc, err := p.Parse()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", sc)

	for _, tbl := range sc.Tables {
		spew.Dump(tbl)
	}

}
