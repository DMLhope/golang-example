package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}
func vals() (int, int) {
	return 3, 7
}
func main() {
	res := plus(vals())
	fmt.Println("3+7 =", res)
}
