package main

import (
	"bufio"
	"fmt"
	"os"
)

func doCat() error {
	argv := os.Args
	argc := len(argv)

	if argc < 2 {
		return fmt.Errorf("ファイルを指定してください")
	}

	for i := 1; i < argc; i++ {
		f, err := os.Open(argv[i])
		defer f.Close()
		if err != nil {
			return fmt.Errorf("ファイルを読み込めませんでした。%s", f)
		}
		sc := bufio.NewScanner(f)
		for sc.Scan() {
			fmt.Print(sc.Text())
		}
		if err := sc.Err(); err != nil {
			return fmt.Errorf("")
		}
	}
	return nil
}

func main() {
	if err := doCat(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
