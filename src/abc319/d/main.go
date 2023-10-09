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

	// 入力の整形
	n := atoi(strings.Split(lines[0], " ")[0])
	m := atoi(strings.Split(lines[0], " ")[1])

	ls := make([]int, n)
	strs := strings.Split(lines[1], " ")
	for i := 0; i < n; i++ {
		ls[i] = atoi(strs[i]) + 1 // 単語の間のスペースを考慮する
	}

	lower := max(ls) - 1 // 最小の行数(単語の間のスペースを引く)
	upper := sum(ls)     // 最大の行数

	// 2分探索で求める
	for lower+1 < upper {
		mid := (lower + upper) / 2

		row := 1 // 今の行数
		len := 0 // 先頭から何文字目か

		for i := 0; i < n; i++ {
			len += ls[i]   // 行の末尾に追加してみる
			if len > mid { // はみ出たら
				row++       // 改行して
				len = ls[i] // その単語を次の行に追加する
			}
		}

		if row <= m { // 今の行数がm以下なら
			upper = mid // mid以下の行数で可能
		} else {
			lower = mid // mid以上の行数で必要
		}
	}

	fmt.Fprintln(wtr, upper-1) // 単語の間のスペースを引く
}

// ////////////////////////////////////////////////////////////////
// functions
// ///////////////////////////////////////////////////////////////

// getStdin: 標準入力を取得する
func getStdin() []string {
	stdin, _ := io.ReadAll(os.Stdin)
	return strings.Split(strings.TrimRight(string(stdin), "\n"), "\n")
}

// flush stdoutに書き込み
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
