package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-cad"
)

func TestStringer(t *testing.T) {
	c := cad.Cents(204)

	expected := "CAD$2.04"
	result := fmt.Sprintf("%s", c)

	if expected != result {
		t.Fatal(fmt.Sprintf("Test Failed.\nExpected: %s\nGot: %s", expected, result))
	}
}
