// USAGE: cat sample.txt | ./head1 n
package main

import (
	"fmt"
	"os"
)

func do_head() {

}

func run() {
	argv := os.Args
	argc := len(argv)
	if argc != 2 {
		return fmt.Erroff(stderr, "引数の数が間違っています。正しい利用法はこちら: %s n\n", argv[0])
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
