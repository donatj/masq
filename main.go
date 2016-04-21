package main

import (
	"bufio"
	"fmt"
	"os"
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
		fmt.Printf("%#v", tbl.TableName)
	}

}
