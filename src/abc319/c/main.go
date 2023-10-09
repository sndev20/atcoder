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
	c := [9]int{}
	for i := range c {
		fmt.Scan(&c[i])
	}

	o := make([]int, 9)
	for i := range o {
		o[i] = i
	}

	// 縦・横・斜めの列の一覧
	vec := [][]int{
		{0, 1, 2}, // 上から 1 行目
		{3, 4, 5}, // 上から 2 行目
		{6, 7, 8}, // 上から 3 行目
		{0, 3, 6}, // 左から 1 列目
		{1, 4, 7}, // 左から 2 列目
		{2, 5, 8}, // 左から 3 列目
		{0, 4, 8}, // 左上から右下
		{2, 4, 6}, // 右上から左下
	}

	var t int
	var r int
	for {
		var disapointed bool
		for _, v := range vec {
			if c[v[0]] == c[v[1]] && o[v[0]] < o[v[2]] && o[v[1]] < o[v[2]] {
				disapointed = true
				break
			} else if c[v[1]] == c[v[2]] && o[v[1]] < o[v[0]] && o[v[2]] < o[v[0]] {
				disapointed = true
				break
			} else if c[v[2]] == c[v[0]] && o[v[2]] < o[v[1]] && o[v[0]] < o[v[1]] {
				disapointed = true
				break
			}
		}
		if !disapointed { // がっかりしない
			r++
		}
		t++ // 総数
		if !nextPermutation(sort.IntSlice(o)) {
			break
		}
	}
	fmt.Fprintf(wtr, "%.9f\n", (float64(r) / float64(t)))
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
