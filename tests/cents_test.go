package tests

import (
	"testing"

	"github.com/aliforever/go-cad"
)

func TestCents(t *testing.T) {
	data := []struct {
		Cents int64
	}{
		{
			Cents: 250,
		},
		{
			Cents: 305,
		},
		{
			Cents: -650,
		},
		{
			Cents: -400,
		},
	}

	for testNumber, datum := range data {
		result := cad.Cents(datum.Cents).AsCents()
		if result != datum.Cents {
			t.Errorf("For test #%d: Test Failed.", testNumber)
			t.Logf("EXPECTED: %d", datum.Cents)
			t.Logf("ACTUAL:   %d", result)
			continue
		}
	}
}
