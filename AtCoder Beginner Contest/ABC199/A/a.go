package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := newScanner()
	a, b, c := sc.readInt(), sc.readInt(), sc.readInt()
	if a*a+b*b < c*c {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

type scanner struct {
	bufScanner *bufio.Scanner
}

func newScanner() *scanner {
	bufSc := bufio.NewScanner(os.Stdin)
	bufSc.Split(bufio.ScanWords)
	bufSc.Buffer(nil, 100000000)
	return &scanner{bufSc}
}

func (sc *scanner) readString() string {
	bufSc := sc.bufScanner
	bufSc.Scan()
	return bufSc.Text()
}

func (sc *scanner) readInt() int {
	num, err := strconv.Atoi(sc.readString())
	if err != nil {
		panic(err)
	}
	return num
}
