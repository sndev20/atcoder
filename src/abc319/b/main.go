package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var wtr = bufio.NewWriter(os.Stdout)

func main() {
	defer flush()

	var n int
	fmt.Scan(&n)

	ans := ""
	for i := 0; i <= n; i++ {
		for j := 1; j <= 9; j++ {
			if n%j == 0 && i%(n/j) == 0 {
				ans += itoa(j)
				break
			}
		}
		if len(ans) < i+1 {
			ans += "-"
		}
	}

	fmt.Fprintln(wtr, ans)
}

// ////////////////////////////////////////////////////////////////
// functions
// ///////////////////////////////////////////////////////////////

// getStdin: 標準入力を取得する
func getStdin() []string {
	stdin, _ := io.ReadAll(os.Stdin)
	return strings.Split(strings.TrimRight(string(stdin), "\n"), "\n")
}

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

// itoa: intをstringに変換する
func itoa(n int) string {
	return strconv.Itoa(n)
}
