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
	"sync"
)

var(
	arg0 string
	addr = "127.0.0.1"
	stdtp = "tcp"
	wr *csv.Writer
	wg sync.WaitGroup
)

func
ScanPort(rec []string) {
	tp := rec[0]
	a := rec[1]

	conn, err := net.Dial(tp, a)
	if err == nil {
		conn.Close()
	}

	t := time.Now()
	rec = append(
		[]string{
			t.String(),
		},
		rec...,
	)
	if err != nil {
		rec = append(rec, err.Error())
	} else {
		rec = append(rec, "")
	}

	wr.Write(rec)
	wr.Flush()

	wg.Done()
}

func
Run(args []string) {
	arg0 = args[0]
	args = args[1:]
	rd := csv.NewReader(os.Stdin)
	wr = csv.NewWriter(os.Stdout)
	wr.Comma, rd.Comma = '\t', '\t'
	wr.Write([]string{"Time", "Type", "Target", "Error"})
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
			break
		} else if err != nil {
			log.Println(err)
		}
		wg.Add(1)
		go ScanPort(rec)
	}
	wg.Wait()
}

