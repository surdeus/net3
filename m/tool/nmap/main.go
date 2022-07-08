package nmap

import(
	"os"
	"fmt"
	"io"
	"bufio"
)

var(
	rd = bufio.NewReader(os.Stdin)
	arg0 string
	addr = "127.0.0.1"
	tp = "tcp"
)

func
Run(args []string) {
	arg0 = args[0]
	for {
		line, err := rd.ReadString('\n')
		fmt.Println(line)
		if err == io.EOF {
			os.Exit(0)
		}
	}
}

