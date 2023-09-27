package jsonstr_test

import (
	"testing"

	"encoding/json"

	"sourcecode.social/reiver/go-jsonstr"
)

func TestString_MarshalJSON(t *testing.T) {

	tests := []struct{
		Strings jsonstr.Strings
		Expected string
	}{
		{
			Strings: jsonstr.Strings{},
			Expected: `[]`,
		},



		{
			Strings: jsonstr.CompileStrings("apple"),
			Expected: `["apple"]`,
		},
		{
			Strings: jsonstr.CompileStrings("apple", "banana"),
			Expected: `["apple","banana"]`,
		},
		{
			Strings: jsonstr.CompileStrings("apple", "banana", "cherry"),
			Expected: `["apple","banana","cherry"]`,
		},



		{
			Strings: jsonstr.CompileStrings("😈"),
			Expected: `["😈"]`,
		},
		{
			Strings: jsonstr.CompileStrings("😈", "🙂🙁"),
			Expected: `["😈","🙂🙁"]`,
		},
		{
			Strings: jsonstr.CompileStrings("😈", "🙂🙁", ""),
			Expected: `["😈","🙂🙁",""]`,
		},
		{
			Strings: jsonstr.CompileStrings("😈", "🙂🙁", "", "٠١٢٣۴۵۶٧٨٩"),
			Expected: `["😈","🙂🙁","","٠١٢٣۴۵۶٧٨٩"]`,
		},
	}

	for testNumber, test := range tests {

		actualBytes, err := json.Marshal(test.Strings)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("STRINGS: %#v", test.Strings)
			t.Logf("EXPECTED: %#v", test.Expected)
			continue
		}

		{
			actual := string(actualBytes)
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual marshaled value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %s", expected)
				t.Logf("ACTUAL:   %s", actual)
				t.Logf("STRINGS: %#v", test.Strings)
				continue
			}
		}
	}
}
