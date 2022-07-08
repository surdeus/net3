package nmap

import(
	"os"
	"io"
	"net"
	"encoding/csv"
	"errors"
	"log"
	"time"
	//"fmt"
)

var(
	arg0 string
	addr = "127.0.0.1"
	stdtp = "tcp"
	wr *csv.Writer
)

func
ScanPort(rec []string) error {

	tp := rec[0]
	a := rec[1]
	_, err := net.Dial(tp, a)
	if err != nil {
		log.Println(err)
		return err
	}

	t := time.Now()
	rec = append([]string{t.String()}, rec...)
	wr.Write(rec)
	wr.Flush()

	return nil
}

func
Run(args []string) {
	arg0 = args[0]
	args = args[1:]
	rd := csv.NewReader(os.Stdin)
	wr = csv.NewWriter(os.Stdout)
	wr.Comma,  rd.Comma = '\t', '\t'
	wr.Write([]string{"Time", "Type", "Target"})
	wr.Flush()
	rec, err := rd.Read()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	} else if len(rec) < 2 {
		err := errors.New("Wrong CSV format: must be at least 2 fields")
		log.Println(err)
		os.Exit(1)
	}
	for {
		rec, err := rd.Read()
		if err == io.EOF {
			os.Exit(0)
		} else if err != nil {
			log.Println(err)
		}
		go ScanPort(rec)
	}
}

