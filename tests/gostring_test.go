package tests

import (
	"testing"

	"github.com/aliforever/go-cad"
)

func TestGoString(t *testing.T) {
	data := []struct {
		Cents    int64
		Expected string
	}{
		{
			Cents:    224,
			Expected: "cad.Cents(224)",
		},
		{
			Cents:    310,
			Expected: "cad.Cents(310)",
		},
		{
			Cents:    156,
			Expected: "cad.Cents(156)",
		},
		{
			Cents:    -230,
			Expected: "cad.Cents(-230)",
		},
	}
	for testNumber, datum := range data {
		result := cad.Cents(datum.Cents).GoString()
		if result != datum.Expected {
			t.Errorf("For test #%d: Test Failed.", testNumber)
			t.Logf("EXPECTED: %s", datum.Expected)
			t.Logf("ACTUAL:   %s", result)
			continue
		}
	}
}

func TestParseCadGoString(t *testing.T) {
	data := []struct {
		CAD      string
		Expected string
	}{
		{
			CAD:      "$2.24",
			Expected: "cad.Cents(224)",
		},
		{
			CAD:      "$.95",
			Expected: "cad.Cents(95)",
		},
		{
			CAD:      "$-1.56",
			Expected: "cad.Cents(-156)",
		},
		{
			CAD:      "9¢",
			Expected: "cad.Cents(9)",
		},
	}
	for testNumber, datum := range data {
		c, err := cad.ParseCAD(datum.CAD)
		if err != nil {
			t.Errorf("Can't parse cad for %s => %s", datum.CAD, err)
			continue
		}
		result := c.GoString()
		if result != datum.Expected {
			t.Errorf("For test #%d: Test Failed.", testNumber)
			t.Logf("EXPECTED: %s", datum.Expected)
			t.Logf("ACTUAL:   %s", result)
			continue
		}
	}
}
