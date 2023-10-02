package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var wtr = bufio.NewWriter(os.Stdout)

func init() {
}

func main() {
	defer flush()

	// 入力値の受け取り
	var n, k int
	var s string
	fmt.Scanf("%d %d", &n, &k)
	fmt.Scanf("%s", &s)

	ans := 0
	if n-k > 0 {
		ans++
		n -= k
		ans += n / (k - 1)
	}
	if n%(k-1) > 0 {
		ans++
	}

	fmt.Fprintln(wtr, ans)
}

// ////////////////////////////////////////////////////////////////
// functions
// ///////////////////////////////////////////////////////////////

// flush stdoutを初期化する
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
