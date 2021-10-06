package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-cad"
)

func TestCanonical(t *testing.T) {
	data := []struct {
		Cents    int64
		Expected string
	}{
		{
			Cents:    105,
			Expected: "1 5",
		},
		{
			Cents:    250,
			Expected: "2 50",
		},
		{
			Cents:    -300,
			Expected: "-3 0",
		},
		{
			Cents:    -250,
			Expected: "-2 -50",
		},
	}

	for testNumber, datum := range data {
		dollars, cents := cad.Cents(datum.Cents).CanonicalForm()
		result := fmt.Sprintf("%d %d", dollars, cents)
		if result != datum.Expected {
			t.Errorf("For test #%d: Test Failed.", testNumber)
			t.Logf("EXPECTED: %s", datum.Expected)
			t.Logf("ACTUAL:   %s", result)
			continue
		}
	}
}

func TestCanonicalParse(t *testing.T) {
	data := []struct {
		CAD      string
		Expected string
	}{
		{
			CAD:      "$1.05",
			Expected: "1 5",
		},
		{
			CAD:      "CAD$2.50",
			Expected: "2 50",
		},
		{
			CAD:      "-$3.00",
			Expected: "-3 0",
		},
		{
			CAD:      "$-2.50",
			Expected: "-2 -50",
		},
	}

	for testNumber, datum := range data {
		c, err := cad.ParseCAD(datum.CAD)
		if err != nil {
			t.Errorf("For test #%d: Test Failed. %s", testNumber, err)
			continue
		}
		dollars, cents := c.CanonicalForm()
		result := fmt.Sprintf("%d %d", dollars, cents)
		if result != datum.Expected {
			t.Errorf("For test #%d: Test Failed.", testNumber)
			t.Logf("EXPECTED: %s", datum.Expected)
			t.Logf("ACTUAL:   %s", result)
			continue
		}
	}
}
