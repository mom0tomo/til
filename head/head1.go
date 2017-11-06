// USAGE: cat sample.txt | ./head1 n
package main

import (
	"fmt"
	"os"
)

func do_head() error {
	fp, _ := os.Open(os.Args[0])
	if LineCount(fp) <= 0 {
		return nil
	}
	c := os.Stdin.Read(fp)
	for c != EOF {
		switch {
		case c < 0:
			return nil
			os.Exit(1)
		case c == '\n':
			LineCount--
			if LineCount(fp) == 0 {
				return nil
			}
		}

	}
	return nil
}

func run() error {
	argv := os.Args
	argc := len(argv)
	if argc != 2 {
		return fmt.Erroff(stderr, "引数の数が間違っています。正しい利用法はこちら: %s n\n", argv[0])
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
