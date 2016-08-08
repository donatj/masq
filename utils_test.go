package masq

import "testing"

func TestStrIntSuffixSplit(t *testing.T) {
	type strIntTest struct {
		base string
		s    string
		i    int
		err  error
	}
	tests := []strIntTest{
		{"int11", "int", 11, nil},
		{"float 9", "float ", 9, nil},
		{"float 9 bbq", "", 0, errUnexpectedStringSuffix},
		{"int11 ", "", 0, errUnexpectedStringSuffix},

		{"int", "int", 0, nil},
		{"11", "", 11, nil},
	}
	for _, tt := range tests {
		s, i, err := strIntSuffixSplit(tt.base)
		if tt.err != err {
			t.Fatalf("expected '%#v'; got '%#v'", tt.err, err)
		}
		if tt.s != s {
			t.Fatalf("expected '%s'; got '%s'", tt.s, s)
		}
		if tt.i != i {
			t.Fatalf("expected %d; got %d", tt.i, i)
		}
	}

}
