package tests

import (
	"encoding/json"
	"testing"

	"github.com/aliforever/go-cad"
)

func TestJsonMarshal(t *testing.T) {
	type Data struct {
		Name     string  `json:"name"`
		Balance  cad.CAD `json:"balance"`
		Expected string  `json:"-"`
	}

	d := []Data{
		{
			Name:     "Jow Blow",
			Balance:  cad.Cents(12345),
			Expected: `{"name":"Jow Blow","balance":"CAD$123.45"}`,
		},
		{
			Name:     "John Cena",
			Balance:  cad.Cents(1234),
			Expected: `{"name":"John Cena","balance":"CAD$12.34"}`,
		},
		{
			Name:     "Jow Blow",
			Balance:  cad.Cents(-2452),
			Expected: `{"name":"Jow Blow","balance":"CAD$-24.52"}`,
		},
	}

	for testNumber, data := range d {
		j, err := json.Marshal(data)
		if err != nil {
			t.Errorf("Can't marshal data for test number #%d: %s", testNumber, err)
			continue
		}
		if string(j) != data.Expected {
			t.Errorf("For test #%d: Test Failed.", testNumber)
			t.Logf("EXPECTED: %s", data.Expected)
			t.Logf("ACTUAL:   %s", string(j))
			continue
		}
	}
}
