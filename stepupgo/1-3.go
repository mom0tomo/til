package main

import "fmt"

func main() {

	for i := 1; i < 101; i++ {
		switch i % 2 {
		case 0:
			fmt.Println(i, "-偶数")
		default:
			fmt.Println(i, "-奇数")
		}

	}

}
