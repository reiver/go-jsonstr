package jsonstr

import (
	"github.com/reiver/go-erorr"
)

const (
	ErrNotJSONArrayOfString = erorr.Error("jsonstr: not JSON array of strings")
)

const (
	errNilReceiver = erorr.Error("jsonstr: nil receiver")
	errNilData     = erorr.Error("jsonstr: nil data")
)
