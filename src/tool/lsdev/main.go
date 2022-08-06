package lsdev

import(
	"fmt"
	"github.com/google/gopacket/pcap"
	"log"
)

var(
	arg0 string
)

func
Run(args []string) {
	arg0 = args[0]
	devs, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}
	
	for _, dev := range devs {
		fmt.Printf("%s\t\"%s\"\t",
			dev.Name, dev.Description)
		for _, addr := range dev.Addresses {
			ones, _ := addr.Netmask.Size()
			fmt.Printf("%s/%d,", addr.IP.String(),
				ones)
		}
		fmt.Println("")
	}
}

