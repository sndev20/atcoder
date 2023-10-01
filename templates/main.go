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
	ans := 0

	var n string
	fmt.Scan(&n)

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
