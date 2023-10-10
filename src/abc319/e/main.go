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
	lines := getStdin()

	var n, x, y int
	fmt.Sscan(lines[0], &n, &x, &y)

	pt := make([][]int, n)
	for i := 0; i < n; i++ {
		pt[i] = make([]int, 2)
		fmt.Sscan(lines[i+1], &pt[i][0], &pt[i][1])
	}

	var q int
	fmt.Sscan(lines[n], &q)

	qi := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Sscan(lines[n+1+i], &qi[i])
	}

	tot := make([]int, 840)
	for i := 0; i < 840; i++ {
		now := i
		for j := 0; j < n-1; j++ {
			wait := (pt[j][0] - now%pt[j][0]) % pt[j][0]
			now += wait + pt[j][1]
		}
		// totは、待ち時間の合計＋トータルの乗車時間
		tot[i] = now - i
	}

	for _, v := range qi {
		fmt.Fprintln(wtr, tot[(v+x)%840]+v+x+y)
	}
}

// ////////////////////////////////////////////////////////////////
// functions
// ///////////////////////////////////////////////////////////////

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
