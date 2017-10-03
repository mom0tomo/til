package main

import (
	"fmt"
	"os"
)

func run() error {
	argc := len(os.Args)
	if argc < 2 {
		return fmt.Errorf("引数が足りません。")
	}
	fmt.Println("argc=", argc)

	for i := 0; i < argc; i++ {
		argv := os.Args
		fmt.Printf("argv[%d]=%s\n", i, argv[i])
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
