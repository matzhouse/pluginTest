package processor

import (
	"encoding/json"
	"encoding/xml"
)

// Data is an example data struct we can use to 'process'
// some data
type Data struct {
	ID     string
	TS     string
	Data   []byte
	Active bool
}

// ProcessData takes in a JSON string and outputs an XML string
func ProcessData(in []byte) (out []byte, err error) {

	d := &Data{}

	err = json.Unmarshal(in, d)

	if err != nil {
		return []byte{}, err
	}

	out, err = xml.Marshal(d)

	if err != nil {
		return []byte{}, err
	}

	return out, err

}
