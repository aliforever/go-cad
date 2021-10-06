package tests

import (
	"testing"

	"github.com/aliforever/go-cad"
)

func TestAddCents(t *testing.T) {
	var data = []struct {
		X        int64
		Y        int64
		Expected int64
	}{
		{
			X:        206,
			Y:        309,
			Expected: 515,
		},
		{
			X:        125,
			Y:        652,
			Expected: 777,
		},
		{
			X:        415,
			Y:        216,
			Expected: 631,
		},
		{
			X:        -142,
			Y:        -200,
			Expected: -342,
		},
	}

	for testNumber, d := range data {
		result := cad.Cents(d.X).Add(cad.Cents(d.Y)).AsCents()
		if result != d.Expected {
			t.Errorf("For test #%d: Test Failed.", testNumber)
			t.Logf("EXPECTED: %d", d.Expected)
			t.Logf("ACTUAL:   %d", result)
			continue
		}
	}
}

func TestAddParse(t *testing.T) {
	var data = []struct {
		X        string
		Y        string
		Expected int64
	}{
		{
			X:        "$2.06",
			Y:        "CAD$3.09",
			Expected: 515,
		},
		{
			X:        "125¢",
			Y:        "$6.52",
			Expected: 777,
		},
		{
			X:        "$4.15",
			Y:        "$2.16",
			Expected: 631,
		},
		{
			X:        "-$1.42",
			Y:        "$-2.00",
			Expected: -342,
		},
	}

	for testNumber, d := range data {
		xCad, err := cad.ParseCAD(d.X)
		if err != nil {
			t.Errorf("For test #%d: Test Failed for %s. %s", testNumber, d.X, err)
			continue
		}
		yCad, err := cad.ParseCAD(d.Y)
		if err != nil {
			t.Errorf("For test #%d: Test Failed for %s. %s", testNumber, d.Y, err)
			continue
		}
		result := xCad.Add(yCad).AsCents()
		if result != d.Expected {
			t.Errorf("For test #%d: Test Failed.", testNumber)
			t.Logf("EXPECTED: %d", d.Expected)
			t.Logf("ACTUAL:   %d", result)
			continue
		}
	}
}
