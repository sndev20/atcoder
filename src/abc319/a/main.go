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
	lines := getStdin()

	rating := map[string]int{
		"tourist":    3858,
		"ksun48":     3679,
		"Benq":       3658,
		"Um_nik":     3648,
		"apiad":      3638,
		"Stonefeang": 3630,
		"ecnerwala":  3613,
		"mnbvmar":    3555,
		"newbiedmy":  3516,
		"semiexp":    3481,
	}

	fmt.Fprintln(wtr, rating[lines[0]])
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
