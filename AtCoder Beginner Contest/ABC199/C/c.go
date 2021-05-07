package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := newScanner()
	n := sc.readInt()
	s := sc.readRunes()
	q := sc.readInt()
	front, back := s[:n], s[n:]
	for i := 0; i < q; i++ {
		t, a, b := sc.readInt(), sc.readInt(), sc.readInt()
		if t == 1 {
			a--
			b--
			a0 := front
			a1 := front
			if a >= n {
				a0 = back
				a -= n
			}
			if b >= n {
				a1 = back
				b -= n
			}
			a0[a], a1[b] = a1[b], a0[a]
		} else {
			front, back = back, front
		}
	}
	ans := string(front) + string(back)
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

func (sc *scanner) readRunes() []rune {
	return []rune(sc.readString())
}
