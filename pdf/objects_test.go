package pdf

import "bytes"
import "testing"


// First define some helper functions

func toString (object Object) string {
	var buffer bytes.Buffer

	object.Serialize (&buffer)
	return buffer.String()
}

func testOneBoolean (t *testing.T, value bool, expect string) {
	if s := toString(NewBoolean(value)); s != expect {
		t.Errorf (`NewBoolean(%v) produced "%s"; expected "%s"`, value, s, expect)
	}
}

func testOneNumeric (t *testing.T, testvalue float64, expect string) {
	if s := toString(NewNumeric(testvalue)); s != expect {
		t.Errorf (`NewNumeric(%g) produced "%s"; expected "%s"`, testvalue, s, expect)
	}
}

func testOneName (t *testing.T, name, expect string) {
	if s := toString(NewName(name)); s != expect {
		t.Errorf (`NewName(%s) produced "%s"`, name, s)
	}
}

// Unit tests follow

func TestNull(t *testing.T) {
	expect := "null"
	if s := toString(&Null{}); s != expect {
		t.Errorf (`null.Serialize() produced "%s"; expected "%s"`, s, expect)
	}
}


func TestBoolean(t *testing.T) {
	testOneBoolean (t, false, "false")
	testOneBoolean (t, true, "true")
}


func TestNumeric(t *testing.T) {
	testOneNumeric(t, 1, "1")
	testOneNumeric(t, 3.14159, "3.14159")
	testOneNumeric(t, 0.1, "0.1")
	testOneNumeric(t, 2147483647,  "2147483647")
	testOneNumeric(t, -2147483648, "-2147483648")
	testOneNumeric(t, 3.403e+38,   "3.4028235e+38")
	testOneNumeric(t, -3.403e+38,  "-3.4028235e+38")

	// The PDF spec recommends setting anything below +/-
	// 1.175e-38 to 0 in case a conforming reader uses 32 bit
	// floats.  Here, Adobe is referring to the smallest number
	// that can be represented without losing precision rather
	// than the smallest number that can be represented with a
	// float32. It's odd that Adobe thinks that setting small
	// numbers to zero is better than accepting a representable
	// number with a loss of precision.

	testOneNumeric(t, 1.176e-38, "1.176e-38")
	testOneNumeric(t, -1.176e-38, "-1.176e-38")
	testOneNumeric(t, 1.175e-38, "0")
	testOneNumeric(t, -1.175e-38, "0")
}

func TestOneName (t *testing.T) {
	testOneName (t, "foo", "/foo")
	testOneName (t, "résumé", "/résumé")
	testOneName (t, "foo bar", "/foo#20bar")
}

