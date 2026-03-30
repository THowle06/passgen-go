package generator

import "testing"

func TestLength(t *testing.T) {
	policy := Policy{Length: 16, IncludeUpper: true, IncludeLower: true, IncludeDigits: true, IncludeSymbols: true}
	pwd := Generate(policy)

	if len(pwd) != 16 {
		t.Errorf("expected length 16, got %d", len(pwd))
	}
}
