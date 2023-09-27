package jsonstr

import (
	"encoding/json"
)

func validate(data []byte) (bool, error) {
	if nil == data {
		return false, errNilData
	}

	var ss []string

	err := json.Unmarshal(data, &ss)
	if nil != err {
		switch err.(type) {
		case *json.UnmarshalTypeError:
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
