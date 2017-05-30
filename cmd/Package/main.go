package main

import (
	"github.com/matzhouse/plugintest/pkg/processor"
)

func ProcessData(in []byte) (out []byte, err error) {

	return processor.ProcessData(in)

}
