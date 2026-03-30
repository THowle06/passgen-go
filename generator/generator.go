package generator

import (
	"crypto/rand"
	"math"
	"math/big"
)

type Policy struct {
	Length           int
	IncludeUpper     bool
	IncludeLower     bool
	IncludeDigits    bool
	IncludeSymbols   bool
	RequireEachClass bool
}

var upper = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var lower = []rune("abcdefghijklmnopqrstuvwxyz")
var digits = []rune("0123456789")
var symbols = []rune("!@#$%^&*()-_+=?")

func randomChar(set []rune) rune {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(set))))
	return set[n.Int64()]
}

func shuffle(runes []rune) {
	for i := len(runes) - 1; i > 0; i-- {
		jBig, _ := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		j := int(jBig.Int64())
		runes[i], runes[j] = runes[j], runes[i]
	}
}

func Generate(policy Policy) string {
	var all []rune
	var password []rune
	var sets [][]rune

	if policy.IncludeUpper {
		all = append(all, upper...)
		sets = append(sets, upper)
	}
	if policy.IncludeLower {
		all = append(all, lower...)
		sets = append(sets, lower)
	}
	if policy.IncludeDigits {
		all = append(all, digits...)
		sets = append(sets, digits)
	}
	if policy.IncludeSymbols {
		all = append(all, symbols...)
		sets = append(sets, symbols)
	}

	// Require each class
	if policy.RequireEachClass {
		for _, set := range sets {
			password = append(password, randomChar(set))
		}
	}

	// Fill remaining
	for len(password) < policy.Length {
		password = append(password, randomChar(all))
	}

	shuffle(password)

	return string(password)
}

func CalculateEntropy(policy Policy) float64 {
	charsetSize := 0

	if policy.IncludeUpper {
		charsetSize += 26
	}
	if policy.IncludeLower {
		charsetSize += 26
	}
	if policy.IncludeDigits {
		charsetSize += 10
	}
	if policy.IncludeSymbols {
		charsetSize += len(symbols)
	}

	return float64(policy.Length) * math.Log2(float64(charsetSize))
}

func ClassifyEntropy(entropy float64) string {
	if entropy < 40 {
		return "Weak"
	} else if entropy < 80 {
		return "Moderate"
	}
	return "Strong"
}
