package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc := newScanner()
	n, m := sc.readInt(), sc.readInt()
	max := int(math.Pow(2, float64(n)))
	dp := make([]int, max)
	check := make([]bool, max)
	constraints := make([][]pair, n+1)
	for i := 0; i < m; i++ {
		x, y, z := sc.readInt(), sc.readInt()-1, sc.readInt()
		constraints[x] = append(constraints[x], pair{y, z})
	}
	for bit := 0; bit < max; bit++ {
		nums := make([]int, n)
		count := 0
		for j := 0; j < n; j++ {
			if ((bit >> j) & 1) == 0 {
				continue
			}
			nums[j]++
			count++
		}
		for i := 0; i < n-1; i++ {
			nums[i+1] += nums[i]
		}
		for _, p := range constraints[count] {
			if nums[p.first] > p.second {
				check[bit] = true
				break
			}
		}
	}
	dp[0] = 1

	for bit := 0; bit < max; bit++ {
		if check[bit] {
			continue
		}
		for i := 0; i < n; i++ {
			if ((bit >> i) & 1) == 0 {
				continue
			}
			prev := bit &^ (1 << i)
			dp[bit] += dp[prev]
		}
	}
	Print(dp[max-1])
}

type pair struct {
	first  int
	second int
}

func Print(values ...interface{}) {
	for _, v := range values {
		fmt.Printf("%v ", v)
	}
	fmt.Println("")
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
