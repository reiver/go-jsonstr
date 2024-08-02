package jsonstr_test

import (
	"testing"

	"encoding/json"

	"github.com/reiver/go-jsonstr"
)

func TestString_UnmarshalJSON(t *testing.T) {

	tests := []struct{
		JSON string
		Expected jsonstr.Strings
	}{
		{
			JSON: `[]`,
			Expected: jsonstr.Strings{},
		},
		{
			JSON: `[ ]`,
			Expected: jsonstr.Strings{},
		},
		{
			JSON: `[  ]`,
			Expected: jsonstr.Strings{},
		},



		{
			JSON:                          `["apple"]`,
			Expected: jsonstr.CompileStrings("apple"),
		},
		{
			JSON:                          `["apple","banana"]`,
			Expected: jsonstr.CompileStrings("apple","banana"),
		},
		{
			JSON:                          `["apple","banana","cherry"]`,
			Expected: jsonstr.CompileStrings("apple","banana","cherry"),
		},



		{
			JSON:                          `["😈"]`,
			Expected: jsonstr.CompileStrings("😈"),
		},
		{
			JSON:                          `["😈","🙂🙁"]`,
			Expected: jsonstr.CompileStrings("😈","🙂🙁"),
		},
		{
			JSON:                          `["😈","🙂🙁",""]`,
			Expected: jsonstr.CompileStrings("😈","🙂🙁",""),
		},
		{
			JSON:                          `["😈","🙂🙁","","٠١٢٣۴۵۶٧٨٩"]`,
			Expected: jsonstr.CompileStrings("😈","🙂🙁","","٠١٢٣۴۵۶٧٨٩"),
		},



		{
			JSON:                          `["1","two",  "THREE",  "iv", "۵"]`,
			Expected: jsonstr.CompileStrings("1","two",  "THREE",  "iv", "۵"),
		},
	}

	for testNumber, test := range tests {

		var actual jsonstr.Strings

		err := json.Unmarshal([]byte(test.JSON), &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("JSON: %#v", test.JSON)
			t.Logf("EXPECTED: %#v", test.Expected)
			continue
		}

		{
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual marshaled value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("JSON: %#v", test.JSON)
				continue
			}
		}
	}
}

func TestString_UnmarshalJSON_fail(t *testing.T) {

	tests := []struct{
		JSON string
	}{
		{
			JSON: `{}`,
		},
		{
			JSON: `{"name":"value"}`,
		},



		{
			JSON: `false`,
		},
		{
			JSON: `true`,
		},



		{
			JSON: `-2.223`,
		},
		{
			JSON: `-1`,
		},
		{
			JSON: `0`,
		},
		{
			JSON: `11`,
		},
		{
			JSON: `222.22`,
		},
		{
			JSON: `3333`,
		},
		{
			JSON: `44444`,
		},
		{
			JSON: `5.0`,
		},



		{
			JSON: `[false]`,
		},
		{
			JSON: `[false, true]`,
		},
		{
			JSON: `[false, false, false, false, true, true, true, false]`,
		},



		{
			JSON: `["hello", 1, "wow", true, -3.2, "world"]`,
		},

		{
			JSON: `["once","twice","thrice","fource",5]`,
		},
	}

	for testNumber, test := range tests {

		var actual jsonstr.Strings

		err := json.Unmarshal([]byte(test.JSON), &actual)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("JSON: %#v", test.JSON)
			continue
		}

		{
			expected := jsonstr.ErrNotJSONArrayOfString
			actual   := err

			if expected != actual {
				t.Errorf("For test #%d, the actual error is not what was expected.", testNumber)
				t.Logf("EXPECTED ERROR: (%T) %s", expected, expected)
				t.Logf("ACTUAL   ERROR: (%T) %s", actual, actual)
				t.Logf("JSON: %#v", test.JSON)
				continue
			}
		}
	}
}
