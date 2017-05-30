package main

import (
	"fmt"
	"time"

	"github.com/matzhouse/plugintest/pkg/processor"
)

func main() {

	var dataIn = ""

	var out []byte
	var err error

	t := time.Now()

	for i := 0; i < 10000; i++ {
		out, err = processor.ProcessData([]byte{dataIn})
	}

	n := time.Since(t)

	fmt.Print("10000 loops took ", n)

}
