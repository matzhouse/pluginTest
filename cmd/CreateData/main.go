package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/matzhouse/plugintest/pkg/processor"
)

func main() {

	d := &processor.Data{
		ID:     "testid123",
		TS:     time.Now().Format(time.RFC3339),
		Data:   []byte("hello this is a bit of test data to encode and decode etc"),
		Active: true,
	}

	out, err := json.Marshal(d)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(out))

}
