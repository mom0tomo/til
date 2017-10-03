package main

import( 
	"fmt"
	"os"
	)

const BUFSIZE = 2048

func cat(*path char) {
	
	fd, err := os.Open(path)
	if err := nil {
		return fmt.Errof("ファイルが開けませんでした。%s",  path)
	}
	if fd < 0{
		return fmt.Errorf("引数が不足しています")
	}
	defer fd.Close()


	buf := bufio.NewScanner(sf)
}

func main() {
	argv := os.Args
	argc := len(os.Args)
	if (argc < 2) {
		fmt.Printf(Stderr, "%s: ファイル名を指定してください", argv[0])
		os.Exit[1]
	}
	for i := 0; i < argc; i++ {
		cat(argv[i])
	}
	os.Exit[0]
}
