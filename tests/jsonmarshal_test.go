package tests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/aliforever/go-cad"
)

func TestJsonMarshal(t *testing.T) {
	type Data struct {
		Name    string  `json:"name"`
		Balance cad.CAD `json:"balance"`
	}

	d := Data{Name: "Jow Blow", Balance: cad.Cents(12345)}
	jsonBytes, err := json.Marshal(d)
	if err != nil {
		t.Fatal("Test Failed.", err)
		return
	}

	expected := `{"name":"Jow Blow","balance":"CAD$123.45"}`
	result := string(jsonBytes)

	if expected != result {
		t.Fatal(fmt.Sprintf("Test Failed.\nExpected: %s\nGot: %s", expected, result))
	}
}
