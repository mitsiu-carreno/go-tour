package stringutil

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in, expect string
	}{
		{"Te amo corazón", "nózaroc oma eT"},
		{"Juntos por siempre", "erpmeis rop sotnuJ"},
		{"Quiero verte", "etrev oreiuQ"},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.expect {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.expect)
		}
	}
}
