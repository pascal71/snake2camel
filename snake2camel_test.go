package snake2camel

import "testing"

func TestToCamelCase(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"hello_world", "helloWorld"},
		{"this_is_a_test", "thisIsATest"},
		{"example_string", "exampleString"},
		{"", ""},
	}

	for _, c := range cases {
		got := ToCamelCase(c.in)
		if got != c.want {
			t.Errorf("ToCamelCase(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
