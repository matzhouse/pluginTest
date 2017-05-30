package main

import (
	"fmt"
	"time"

	"github.com/matzhouse/plugintest/pkg/processor"
)

func main() {

	var dataRaw = `{"ID":"testid123","TS":"2017-05-30T15:09:07+01:00","Data":"aGVsbG8gdGhpcyBpcyBhIGJpdCBvZi
B0ZXN0IGRhdGEgdG8gZW5jb2RlIGFuZCBkZWNvZGUgZXRj","Active":true}`

	var dataIn = []byte(dataRaw)

	var out []byte
	var err error

	t := time.Now()

	for i := 0; i < 1000000; i++ {
		out, err = processor.ProcessData(dataIn)
	}

	_ = out
	_ = err

	n := time.Since(t)

	fmt.Print("1000000 loops took ", n)

}
