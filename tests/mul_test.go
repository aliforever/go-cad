package tests

import (
	"testing"

	"github.com/aliforever/go-cad"
)

func TestMulCents(t *testing.T) {
	data := []struct {
		Cents      int64
		MultiplyBy int64
		Expected   int64
	}{
		{
			MultiplyBy: 2,
			Cents:      106,
			Expected:   212,
		},
		{
			MultiplyBy: 3,
			Cents:      -150,
			Expected:   -450,
		},
	}

	for testNumber, datum := range data {
		result := cad.Cents(datum.Cents).Mul(datum.MultiplyBy).AsCents()

		if result != datum.Expected {
			t.Errorf("For test #%d: Test Failed.", testNumber)
			t.Logf("EXPECTED: %d", datum.Expected)
			t.Logf("ACTUAL:   %d", result)
			continue
		}
	}
}

func TestMulParse(t *testing.T) {
	data := []struct {
		CAD        string
		MultiplyBy int64
		Expected   int64
	}{
		{
			MultiplyBy: 2,
			CAD:        "$1.06",
			Expected:   212,
		},
		{
			MultiplyBy: 3,
			CAD:        "-$1.50",
			Expected:   -450,
		},
	}

	for testNumber, datum := range data {
		c, err := cad.ParseCAD(datum.CAD)
		if err != nil {
			t.Errorf("Can't parse cad for %s => %s", datum.CAD, err)
			continue
		}

		result := c.Mul(datum.MultiplyBy).AsCents()

		if result != datum.Expected {
			t.Errorf("For test #%d: Test Failed.", testNumber)
			t.Logf("EXPECTED: %d", datum.Expected)
			t.Logf("ACTUAL:   %d", result)
			continue
		}
	}
}
