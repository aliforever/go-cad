package tests

import (
	"testing"

	"github.com/aliforever/go-cad"
)

func TestStringer(t *testing.T) {
	d := []struct {
		Cents    int64
		Expected string
	}{
		{
			Cents:    204,
			Expected: "CAD$2.04",
		},
		{
			Cents:    -150,
			Expected: "CAD$-1.50",
		},
		{
			Cents:    -9,
			Expected: "CAD$-0.09",
		},
		{
			Cents:    90,
			Expected: "CAD$0.90",
		},
	}

	for testNumber, data := range d {
		result := cad.Cents(data.Cents)
		if result.String() != data.Expected {
			t.Errorf("For test #%d: Test Failed.", testNumber)
			t.Logf("EXPECTED: %s", data.Expected)
			t.Logf("ACTUAL:   %s", result.String())
			continue
		}
	}
}

func TestParseCadStringer(t *testing.T) {
	d := []struct {
		CAD      string
		Expected string
	}{
		{
			CAD:      "-9¢",
			Expected: "CAD$-0.09",
		},
		{
			CAD:      "CAD$-1.50",
			Expected: "CAD$-1.50",
		},
		{
			CAD:      "$0.09",
			Expected: "CAD$0.09",
		},
		{
			CAD:      "90¢",
			Expected: "CAD$0.90",
		},
	}

	for testNumber, data := range d {
		c, err := cad.ParseCAD(data.CAD)
		if err != nil {
			t.Errorf("Can't parse cad for %s => %s", data.CAD, err)
			continue
		}
		result := c.String()
		if result != data.Expected {
			t.Errorf("For test #%d: Test Failed.", testNumber)
			t.Logf("EXPECTED: %s", data.Expected)
			t.Logf("ACTUAL:   %s", result)
			continue
		}
	}
}
