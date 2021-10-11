package tests

import (
	"testing"

	"github.com/aliforever/go-cad"
)

func TestParseCad(t *testing.T) {
	mapValueCents := map[string]int64{
		"-$1,234.56":     -123456,
		"-$1234.56":      -123456,
		"$-1234.56":      -123456,
		"$-1,234.56":     -123456,
		"CAD -$1234.56":  -123456,
		"CAD $-1234.56":  -123456,
		"CAD-$1,234.56":  -123456,
		"CAD$-1,234.56":  -123456,
		"$1234.56":       123456,
		"$1.02":          102,
		"$1,234.56":      123456,
		"CAD $1234.56":   123456,
		"CAD $1,234.56":  123456,
		"CAD$1234.56":    123456,
		"CAD$1,234.56":   123456,
		"$0.09":          9,
		"$.09":           9,
		"-$0.09":         -9,
		"-$.09":          -9,
		"$-0.09":         -9,
		"$-.09":          -9,
		"$-9":            -900,
		"CAD $0.09":      9,
		"CAD $.09":       9,
		"CAD -$0.09":     -9,
		"CAD -$.09":      -9,
		"CAD $-0.09":     -9,
		"CAD $-.09":      -9,
		"CAD$0.09":       9,
		"CAD$.09":        9,
		"CAD-$0.09":      -9,
		"CAD-$.09":       -9,
		"CAD$-0.09":      -9,
		"CAD$-.09":       -9,
		"9¢":             9,
		"-9¢":            -9,
		"123456¢":        123456,
		"-123456¢":       -123456,
		"CAD -$1,234.56": -123456,
		"CAD $-1,234.56": -123456,
	}

	for value, cents := range mapValueCents {
		cd, err := cad.ParseCAD(value)
		if err != nil {
			t.Errorf("Can't parse cad for %s => %s", value, err)
			continue
		}
		if cd.AsCents() != cents {
			t.Errorf("Test Failed.\nExpected Cents for %s: %d\nCents: %d", value, cents, cd.AsCents())
			continue
		}
	}
}

func TestParseCadFail(t *testing.T) {
	values := []string{
		"",
		"-",
		"$",
		"¢",
		".",
		",",
		"CAD",
		"$123CAD.45",
		"CADILLAC",
		"cherry",
		"---5",
		"5-",
		"5--",
		"123CAD45",
		"123$45",
		"123¢45",
		"123.45.6",
		"$$$123.45",
		"$123.45$",
		"-$-123.45",
		"$123.-45",
		"$1.120",
	}
	for i, value := range values {
		_, err := cad.ParseCAD(value)
		if err == nil {
			t.Errorf("Test %d Failed\nExpected an error but parsed for %s", i, value)
			continue
		}
	}
}
