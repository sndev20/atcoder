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

	var n string
	fmt.Scan(&n)

	sum := 0
	for i := 0; i < len(n); i++ {
		sum += atoi(n[i : i+1])
	}

	ans := "No"
	if atoi(n)%sum == 0 {
		ans = "Yes"
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
