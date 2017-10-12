package main

import (
	"os"
	"fmt"
	// "unsafe"
)

const bufferSize = 2048

func doCat(path string) error {
	fd, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("ファイルが開けませんでした%s", fd)
	}
	defer fd.Close()

	buf := make([]byte, bufferSize)
	for {
		n, err := fd.Read(buf)
		if err != nil {
			return fmt.Errorf("ファイルが読み込めませんでした%s", n)
		}
		if n < 0 {
			fd.Close()
		}
		if n == 0 {
			break
		}
		// 怪しい
		s, err := fd.Write(buf)
		if s < 0 {
			fd.Close()
		}
		if err != nil {
			return fmt.Errorf("標準出力できませんでした%s", s)
		}
	}
	return nil
}

func run() error {
	argc := len(os.Args)
	argv := os.Args
	if (argc < 2) {
		return fmt.Errorf("%s: ファイルを指定してください%s")
	}
	for i := 1; i < argc; i++ {
		return doCat(argv[i])
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(0)
	}
}