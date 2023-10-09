package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var wtr = bufio.NewWriter(os.Stdout)

func init() {
}

func main() {
	defer flush()

	var k int
	fmt.Scan(&k)

	cnt := 0
	num := 1
	digit := 0
	for cnt < k {
		nums := eachDigit(num)
		if len(nums) == 1 { // 1~9
			fmt.Fprintln(wtr, num)
			cnt++
			num += int(math.Pow(10, float64(digit)))
			continue
		}
		sum := sumOfEachDigit(num)
		if (sum*int(math.Pow(10, float64(digit))) - num) < 0 {
			num += 9 * int(math.Pow(10, float64(digit)))
			digit++
			continue
		}
		fmt.Fprintln(wtr, num)
		num += int(math.Pow(10, float64(digit)))
		cnt++
	}

}

// ////////////////////////////////////////////////////////////////
// functions
// ///////////////////////////////////////////////////////////////

// eachDigit: 各桁の数値を返す
func eachDigit(n int) []int {
	nums := []int{}
	for n > 0 {
		nums = append(nums, n%10)
		n /= 10
	}
	return nums
}

// sumOfEachDigit: 各桁の和を返す
func sumOfEachDigit(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
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
