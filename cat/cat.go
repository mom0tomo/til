package main

import (
	"os"
	"fmt"
	// "unsafe"
)

const bufferSize = 2048

func main() {
	argc := len(os.Args)
	argv := os.Args
	if (argc < 2) {
		fmt.Fprintf(os.Stderr, "%s: ファイルを指定してください\n", argv[0])
	}
	for i := 1; i < argc; i++ {
		doCat(argv[i])
	}
	os.Exit(0)
}

func doCat(path string) {
	fd, err := os.Open(path)
	if err != nil {
		fmt.Errorf("ファイルが開けませんでした%s", fd)
	}
	defer fd.Close()

	buf := make([]byte, bufferSize)
	for {
		// n, err := fd.Read(fd, buf, unsafe.Sizeof(buf))
		n, err := fd.Read(buf)
		if err != nil {
			fmt.Errorf("ファイルが読み込めませんでした%s", n)
		}
		if n < 0 {
			fd.Close()
		}
		if n == 0 {
			break
		}
		// 怪しい
		output, _ := fd.Write(buf)
		if output < 0 {
			fd.Close()
		}
		fmt.Println(string(buf[:n]))
	}
}
