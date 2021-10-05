package tests

import (
	"testing"

	"github.com/aliforever/go-cad"
)

func TestAbs(t *testing.T) {
	tests := []struct {
		Cents    int64
		Expected cad.CAD
	}{
		{
			Cents:    -105,
			Expected: cad.Cents(105),
		},

		{
			Cents:    -5,
			Expected: cad.Cents(5),
		},
		{
			Cents:    -4,
			Expected: cad.Cents(4),
		},
		{
			Cents:    -3,
			Expected: cad.Cents(3),
		},
		{
			Cents:    -2,
			Expected: cad.Cents(2),
		},

		{
			Cents:    -1,
			Expected: cad.Cents(1),
		},

		{
			Cents:    0,
			Expected: cad.Cents(0),
		},
		{
			Cents:    1,
			Expected: cad.Cents(1),
		},
		{
			Cents:    2,
			Expected: cad.Cents(2),
		},
		{
			Cents:    3,
			Expected: cad.Cents(3),
		},
		{
			Cents:    4,
			Expected: cad.Cents(4),
		},
		{
			Cents:    5,
			Expected: cad.Cents(5),
		},

		{
			Cents:    -12345,
			Expected: cad.Cents(12345),
		},
		{
			Cents:    12345,
			Expected: cad.Cents(12345),
		},
	}

	for testNumber, test := range tests {

		c := cad.Cents(test.Cents)

		actual := c.Abs()
		expected := test.Expected

		if actual != expected {
			t.Errorf("For test #%d: Test Failed.", testNumber)
			t.Logf("EXPECTED: %s", expected)
			t.Logf("ACTUAL:   %s", actual)
			continue
		}
	}
}

func TestAbsParse(t *testing.T) {
	tests := []struct {
		CadString     string
		ExpectedCents int64
	}{
		{
			CadString:     "-$1.05",
			ExpectedCents: 105,
		},
		{
			CadString:     "CAD$-3.75",
			ExpectedCents: 375,
		},
		{
			CadString:     "CAD$3.75",
			ExpectedCents: 375,
		},
	}

	for testNumber, test := range tests {
		xCad, err := cad.ParseCAD(test.CadString)
		if err != nil {
			t.Fatalf("For test %s - %s", test.CadString, err)
			return
		}

		result := xCad.Abs().AsCents()
		if result != test.ExpectedCents {
			t.Errorf("For test #%d: Test Failed.", testNumber)
			t.Logf("EXPECTED: %d", test.ExpectedCents)
			t.Logf("ACTUAL:   %d", result)
			continue
		}

	}
}
