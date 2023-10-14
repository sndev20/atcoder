package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

var wtr = bufio.NewWriter(os.Stdout)

func main() {
	defer flush()

	var n int
	fmt.Scan(&n)

	sum := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		sum += a
	}

	fmt.Fprintln(wtr, sum%100)
}

// ////////////////////////////////////////////////////////////////
// functions
// ///////////////////////////////////////////////////////////////

// nextInt: 標準入力の半角スペースで区切られた整数を取得する
func nextInt() int {
	var n int
	fmt.Scan(&n)
	return n
}

// getStdin: 標準入力を取得する
func getStdin() []string {
	stdin, _ := io.ReadAll(os.Stdin)
	return strings.Split(strings.TrimRight(string(stdin), "\n"), "\n")
}

// flush: stdoutに書き込み
func flush() {
	e := wtr.Flush()
	if e != nil {
		panic(e)
	}
}

// atoi: stringをintに変換する
func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

// itoa: intをstringに変換する
func itoa(n int) string {
	return strconv.Itoa(n)
}

// factorial: 階乗を求める
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// nextPermutation: 次の順列が存在するか判定する
func nextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

// max: sliceの最大値を求める
func max(x []int) int {
	m := x[0]
	for _, v := range x {
		if m < v {
			m = v
		}
	}
	return m
}

// sum: sliceの合計値を求める
func sum(x []int) int {
	s := 0
	for _, v := range x {
		s += v
	}
	return s
}
