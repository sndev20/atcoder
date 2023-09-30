package main

import "fmt"

func init() {}

func main() {
	var s string
	fmt.Scan(&s)
	ans := 0
	for i := 0; i < len(s); i++ {
		if s[i:i+1] == "+" {
			ans++
		} else {
			ans--
		}
	}
	fmt.Println(ans)
}
