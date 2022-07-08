package main

import(
	"fmt"
	"os"
	"github.com/k1574/net3/m/tool/lsdev"
	"github.com/k1574/net3/m/tool/nmap"
	"github.com/k1574/net3/m/tool/cap"
)

func main() {
	var(
		utilName string
		args []string
	)
	

	utilsMap := map[string]  interface{}  {
		"map" : nmap.Run,
		"lsdev" : lsdev.Run,
		"cap" : cap.Run,
	}

	if len(os.Args)<2  {
		for k, _ := range utilsMap {
			fmt.Printf("%s\n", k)
		}
		os.Exit(0)
	} else {
		utilName = os.Args[1]
		args = os.Args[1:]
	}

	if _, ok := utilsMap[utilName] ; !ok {
		fmt.Printf("%s: %s: no such util\n", os.Args[0], utilName )
		os.Exit(1)
	}

	utilsMap[utilName].(func([]string) )(args)
}
