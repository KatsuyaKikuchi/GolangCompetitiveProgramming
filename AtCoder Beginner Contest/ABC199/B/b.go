package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(val int, a0 []int, a1 []int) bool {
	n := len(a0)
	for i := 0; i < n; i++ {
		if a0[i] <= val && val <= a1[i] {
			continue
		}
		return false
	}
	return true
}

func main() {
	sc := newScanner()
	n := sc.readInt()
	a0 := sc.readInts(n)
	a1 := sc.readInts(n)

	ans := 0
	for x := 0; x <= 1000; x++ {
		if check(x, a0, a1) {
			ans++
		}
	}

	fmt.Println(ans)
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

func (sc *scanner) readInts(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = sc.readInt()
	}
	return arr
}
