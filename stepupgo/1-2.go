package main

import "fmt"

func main() {

	for i := 1; i < 101; i++ {
		if i%2 == 0 {
			fmt.Println(i, "-偶数")
		} else {
			fmt.Println(i, "-奇数")
		}

	}

}
