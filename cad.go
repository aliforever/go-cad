package cad

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type CAD struct {
	whole      int64
	decimal    int64
	cents      int64
	isNegative bool
}

// Cents returns a CAD that represents ‘n’ cents.
//
// For example, if one was to call:
//
//	cad := Cents(105)
//
// Then ‘cad’ would be: $1.05
func Cents(n int64) CAD {
	var isNegative bool
	if n < 0 {
		isNegative = true
		n = int64(math.Abs(float64(n)))
	}
	whole := n / 100
	decimal := n - (whole * 100)

	return CAD{
		isNegative: isNegative,
		whole:      whole,
		decimal:    decimal,
		cents:      n,
	}
}

// ParseCAD parses the string ‘s’ and return the equivalent CAD.
//
// If ‘s’ does not contain a money amount, then ParseCAD returns an error.
//
// Some example valid strings include:
//
// • -$1234.56
// • $-1234.56
// • -$1,234.56
// • $-1,234.56
// • CAD -$1234.56
// • CAD $-1234.56
// • CAD-$1,234.56
// • CAD$-1,234.56
// • $1234.56
// • $1,234.56
// • CAD $1234.56
// • CAD $1,234.56
// • CAD$1234.56
// • CAD$1,234.56
// • $0.09
// • $.09
// • -$0.09
// • -$.09
// • $-0.09
// • $-.09
// • CAD $0.09
// • CAD $.09
// • CAD -$0.09
// • CAD -$.09
// • CAD $-0.09
// • CAD $-.09
// • CAD$0.09
// • CAD$.09
// • CAD-$0.09
// • CAD-$.09
// • CAD$-0.09
// • CAD$-.09
// • 9¢
// • -9¢
// • 123456¢
// • -123456¢
func ParseCAD(s string) (cad CAD, err error) {
	s = strings.ReplaceAll(s, "$", "")
	s = strings.ReplaceAll(s, "CAD", "")
	s = strings.ReplaceAll(s, "¢", "")
	s = strings.ReplaceAll(s, ",", "")
	s = strings.TrimSpace(s)

	possibleErr := errors.New("invalid_cad")

	isNegative := false
	negativeIndex := strings.Index(s, "-")
	if negativeIndex != -1 && negativeIndex != 0 {
		err = possibleErr
		return
	} else if negativeIndex == 0 {
		isNegative = true
		s = strings.ReplaceAll(s, "-", "")
	}

	if s == "" {
		err = possibleErr
		return
	}

	var whole, decimal int

	split := strings.Split(s, ".")
	if len(split) > 2 {
		err = possibleErr
		return
	}

	if split[0] == "" {
		split[0] = "0"
	}

	if len(split) == 2 {
		if split[1] == "" {
			split[1] = "0"
		}
	}

	var parseErr error

	whole, parseErr = strconv.Atoi(split[0])
	if parseErr != nil {
		fmt.Println(parseErr, split)
		err = possibleErr
		return
	}

	if len(split) == 2 {
		decimal, parseErr = strconv.Atoi(split[1])
		if parseErr != nil {
			err = possibleErr
			return
		}
		if decimal > 99 {
			err = possibleErr
			return
		}
	}

	cents := int64((whole * 100) + decimal)
	if isNegative {
		cents = -cents
	}
	cad = CAD{
		isNegative: isNegative,
		whole:      int64(whole),
		decimal:    int64(decimal),
		cents:      cents,
	}
	return
}

//
// Abs returns the absolute value.
func (c CAD) Abs() CAD {
	return CAD{
		whole:      c.whole,
		decimal:    c.decimal,
		cents:      int64(math.Abs(float64(c.cents))),
		isNegative: false,
	}
}

// AsCents returns CAD as the number of pennies it is equivalent to.
func (c CAD) AsCents() int64 {
	return c.cents
}

// Add adds two CAD and returns the result.
func (c CAD) Add(other CAD) CAD {
	// return Cents((c.whole+other.whole)*100 + (c.cents + other.cents))
	return Cents(c.AsCents() + other.AsCents())
}

// CanonicalForm returns the number of dollars and cents that CAD represents.
//
// ‘cents’ is always less than for equal to 99. I.e.,:
//	cents ≤ 99
func (c CAD) CanonicalForm() (dollars int64, cents int64) {
	whole := c.whole
	if c.isNegative {
		whole = -whole
	}
	return whole, c.decimal
}

// Mul multiplies CAD by a scalar (number) and returns the result.
func (c CAD) Mul(scalar int64) CAD {
	return Cents(c.AsCents() * scalar)
}

// Sub subtracts two CAD and returns the result.
func (c CAD) Sub(other CAD) CAD {
	return Cents(c.AsCents() - other.AsCents())
}
