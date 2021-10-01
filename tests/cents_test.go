package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-cad"
)

func TestCents(t *testing.T) {
	c := cad.Cents(325)
	if c.AsCents() != 325 {
		t.Fatal(fmt.Sprintf("Test Failed. Expected Cents: %d\nCents: %d", 325, c.AsCents()))
		return
	}
}
