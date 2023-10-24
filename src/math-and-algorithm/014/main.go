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
	n := nextInt()

	l := primeFactorization(n)
	// キーだけのスライスを作成
	keys := make([]int, 0, len(l))
	for k := range l {
		keys = append(keys, k)
	}

	// キーをソート
	sort.Ints(keys)

	// 出力
	for _, k := range keys {
		if l[k] != 0 {
			for i := 0; i < l[k]; i++ {
				fmt.Print(k, " ")
			}
		}
	}
	fmt.Println()
}

// ////////////////////////////////////////////////////////////////
// functions
// ///////////////////////////////////////////////////////////////

// primeFactorization: 素因数分解
func primeFactorization(n int) map[int]int {
	res := map[int]int{}
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			res[i]++
			n /= i
		}
	}
	if n != 1 {
		res[n]++
	}
	return res
}

// enumerateDivisors: 約数を列挙する
func enumerateDivisors(n int) []int {
	res := []int{}
	for i := 1; i*i <= n; i++ { // 1から
		if n%i == 0 {
			res = append(res, i)
			if i != n/i {
				res = append(res, n/i)
			}
		}
	}
	sort.Ints(res)
	return res
}

// isPrime: 素数判定
func isPrime(n int) bool {
	// 素数は2 -> √nまで調べて割り切れなければ素数だと言い切れる
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

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
