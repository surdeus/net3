package cap

import(
	"fmt"
	"log"
	"time"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var(
	dev string = "enp3s0"
	snapLen int32 = 1024
	prom bool = false
	timeout time.Duration = 30 * time.Second
	handle *pcap.Handle
	err error
)
	
func
Run(args []string) {
	handle, err = pcap.OpenLive(dev, snapLen, prom, timeout)
	if err != nil {
		log.Fatal(err)
	}

	defer handle.Close()

	pSrc := gopacket.NewPacketSource(handle, handle.LinkType())
	for p := range pSrc.Packets() {
		fmt.Println(p)
	}
}

