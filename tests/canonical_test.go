package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-cad"
)

func TestCanonical(t *testing.T) {
	x := int64(105)

	c := cad.Cents(x)

	dollars, cents := c.CanonicalForm()

	expectedDollars := int64(1)
	expectedCents := int64(5)

	if expectedDollars != dollars {
		t.Fatal(fmt.Sprintf("Test Failed. Expected Dollars: %d\nDollars: %d", expectedDollars, dollars))
		return
	}

	if expectedCents != cents {
		t.Fatal(fmt.Sprintf("Test Failed. Expected Cents: %d\nCents: %d", expectedCents, cents))
		return
	}
}

func TestCanonicalParse(t *testing.T) {
	x := "-$1.05"

	xCad, err := cad.ParseCAD(x)
	if err != nil {
		t.Fatal(err)
		return
	}

	dollars, cents := xCad.CanonicalForm()

	expectedDollars := int64(-1)
	expectedCents := int64(5)

	if expectedDollars != dollars {
		t.Fatal(fmt.Sprintf("Test Failed. Expected Dollars: %d\nDollars: %d", expectedDollars, dollars))
		return
	}

	if expectedCents != cents {
		t.Fatal(fmt.Sprintf("Test Failed. Expected Cents: %d\nCents: %d", expectedCents, cents))
		return
	}
}
