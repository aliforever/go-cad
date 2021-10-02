package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-cad"
)

func TestGoString(t *testing.T) {
	c := cad.Cents(224)

	expected := "Cents(224)"
	result := fmt.Sprintf("%#v", c)

	if expected != result {
		t.Fatal(fmt.Sprintf("Test Failed.\nExpected: %s\nGot: %s", expected, result))
	}
}
