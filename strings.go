package jsonstr

import (
	"bytes"
	"encoding/json"
)

var _ json.Marshaler = Strings{}
var _ json.Unmarshaler = new(Strings)

// Strings represents a JSON array of strings.
//
// For example:
//
// 	var strings jsonstr.Strings
//	
//	// ...
//	
//	jason := []byte(`["once", "twice", "thrice", "fource"]`)
//	
//	err := json.Unmarshal(jason, &strings)
type Strings struct {
	value string
}

func Compile(a ...string) Strings {
	if len(a) <= 0 {
		return Strings{
			value:"[]",
		}
	}

	// This should not return an error, since this is a []string.
	p, _ := json.Marshal(a)

	return Strings{
		value: string(p),
	}
}

func (receiver Strings) Decompile() []string {
	if "" == receiver.value {
		return []string{}
	}

	var ss []string
	json.Unmarshal([]byte(receiver.value), &ss)
	if nil == ss {
		ss = []string{}
	}

	return ss
}

var emptyJSONArray []byte = []byte{'[',']'}

func (receiver Strings) MarshalJSON() ([]byte, error) {
	if "" == receiver.value {
		return emptyJSONArray, nil
	}

	return []byte(receiver.value), nil
}

func (receiver *Strings) UnmarshalJSON(data []byte) error {
	if nil == receiver {
		return errNilReceiver
	}
	if nil == data {
		return errNilData
	}

	var buffer bytes.Buffer
	{

		if err := json.Compact(&buffer, data); nil != err {
			return err
		}
	}

	s := buffer.String()

	if "[]" == s {
		receiver.value = ""
		return nil
	}

	{
		validated, err := validate(buffer.Bytes())
		if nil != err {
			return err
		}
		if !validated {
			return ErrNotJSONArrayOfString
		}
	}

	receiver.value = s
	return nil
}
