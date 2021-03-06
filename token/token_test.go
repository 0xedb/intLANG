package token

import (
	"testing"
)

func TestKeyWords(t *testing.T) {
	tests := []struct {
		name   string
		word   string
		expect bool
	}{
		{"func", "func", true},
		{"string", "string", false},
		{"at", "@", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tok := Token(test.word)

			if got := tok.IsKeyWord(); got != test.expect {
				t.Fatalf("Wanted: %t, Got: %t", test.expect, got)
			}
		})
	}
}

func TestLookupIdentifier(t *testing.T) {
	tests := []struct {
		name   string
		want   string
		expect Token
	}{
		{"At", "@", AT},
		{"Ident", "test", IDENT},
		{"RPAREN", ")", IDENT},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := LookupIdentifier(test.want); got != test.expect {
				t.Fatalf("Wanted: %+v, Got: %+v", test.expect, got)
			}
		})
	}

}
