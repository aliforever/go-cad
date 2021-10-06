package tests

import (
	"encoding/json"
	"testing"

	"github.com/aliforever/go-cad"
)

func TestJsonUnmarshal(t *testing.T) {
	type Data struct {
		Name    string  `json:"name"`
		Balance cad.CAD `json:"balance"`
	}

	data := []struct {
		Json  []byte
		Cents int64
	}{
		{
			Json:  []byte(`{"name":"Jow Blow","balance":"CAD$123.45"}`),
			Cents: 12345,
		},
		{
			Json:  []byte(`{"name":"John Cena","balance":"CAD$-21.01"}`),
			Cents: -2101,
		},
		{
			Json:  []byte(`{"name":"Ali Padidar","balance":"$0.1"}`),
			Cents: 1,
		},
	}

	for testNumber, datum := range data {
		var d *Data
		err := json.Unmarshal(datum.Json, &d)
		if err != nil {
			t.Errorf("Test Failed Unmarshalling for test #%d: %s", testNumber, err)
			continue
		}
		if d.Balance.AsCents() != datum.Cents {
			t.Errorf("For test #%d: Test Failed.", testNumber)
			t.Logf("EXPECTED: %d", datum.Cents)
			t.Logf("ACTUAL:   %d", d.Balance.AsCents())
			continue
		}
	}

}
