package cap

import(
	"os"
	"flag"
	"fmt"
	"log"
	"time"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var(
	arg0 string
	dev string = "enp3s0"
	snapLen int32 = 1024
	prom bool = false
	timeout time.Duration = 30 * time.Second
	handle *pcap.Handle
	err error
)
	
func Run(args []string) {
	arg0 = args[0]
	args = args[1:]

	flagSet := flag.NewFlagSet(arg0, flag.ExitOnError)
	flagSet.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s: %s [options] <dev>\n", arg0, arg0)
		flagSet.PrintDefaults()
		os.Exit(1)
	}
		
	flagSet.Parse(args)
	args = flagSet.Args()
	if len(args) != 1 {
	}

	dev = args[0]

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

