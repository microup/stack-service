package types

import (
	"errors"
	"fmt"
	"net/http"
)

type (
	ResultCode   uint32
	ResultStatus uint32
)

var ErrFileIsNil = errors.New("file is nil")

const (
	ResultOK ResultCode = iota
	ErrLifo
	ErrEmptyValue
	ErrCantFindFile
	ErrCantMarshalJSON
	ErrGotPanic
	StatusBadRequest          ResultCode = http.StatusBadRequest
	StatusInternalServerError ResultCode = http.StatusInternalServerError
	StatusNotExtended         ResultCode = http.StatusNotExtended
	ResultUnknown             ResultCode = 0xFFFFFFFF
)

// ToStr returns a string representation of the ResultCode.
func (s ResultCode) ToStr() string {
	mapStatus := map[ResultCode]string{
		ResultOK:                  "SUCCESS",
		ErrLifo:                   "vlifo error",
		ErrCantMarshalJSON:        "can't marshal json object",
		StatusBadRequest:          "status bad request",
		StatusInternalServerError: "status internal server error",
		StatusNotExtended:         "status not extended",
		ErrEmptyValue:             "value is empty",
		ErrCantFindFile:           "can't find file",
		ErrGotPanic:               "got panic",
		ResultUnknown:             "unknown error",
	}

	m, ok := mapStatus[s]
	if !ok {
		return fmt.Sprintf("unknown result code: %d", s)
	}

	return m
}

// ToUint converts the ResultCode to its uint32 representation.
func (s ResultCode) ToUint() uint32 {
	return uint32(s)
}
