package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mxk/go-flowrate/flowrate"
)

func main() {
	st := time.Now()
	resp, err := http.Get("http://172.16.38.171:2281/livechunksboss/live_110000339_2517179-2020-07-30-18:05:00.flv")
	if err != nil {
		log.Fatalf("Get failed: %v", err)
	}
	defer resp.Body.Close()

	// Limit to 10 bytes per second
	wrappedIn := flowrate.NewReader(resp.Body, -1)

	var f *os.File
	filename := "/Users/lxkaka/Desktop/test_limit.flv"

	if f, err = os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm); err != nil {
		log.Fatalf("openfile failed,filename=%s,err=%+v", filename, err)
		return
	}

	defer f.Close()

	// Copy to stdout
	_, err = io.Copy(f, wrappedIn)
	if err != nil {
		log.Fatalf("Copy failed: %v", err)
	}
	print(time.Since(st).Seconds())
}
