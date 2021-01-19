package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	st := time.Now()
	log.Printf("start:", st)
	c := &http.Client{
		Timeout: 1800 * time.Second,
	}
	resp, err := c.Get("http://upos-sz-office.bilibili.co/livechunks/live_103701980_2364067-2020-10-07-21:45:28.flv?deadline=1602832673&gen=record2vod&os=upos&uparams=deadline,gen,os&upsig=72c415950918262c57e77e234bff4fd8")
	if err != nil {
		log.Fatalf("Get failed: %v", err)
	}
	defer resp.Body.Close()

	// Limit to 10 bytes per second
	//wrappedIn := flowrate.NewReader(resp.Body, -1)

	var f *os.File
	filename := "/Users/lxkaka/Desktop/test_limit.flv"

	if f, err = os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm); err != nil {
		log.Fatalf("openfile failed,filename=%s,err=%+v", filename, err)
		return
	}

	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		println(time.Since(st).Seconds())
		log.Fatalf("Copy failed: %v", err)
	}
	print(time.Since(st).Seconds())
}
