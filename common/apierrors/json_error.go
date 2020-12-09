package apierrors

import (
	"io"

	"github.com/yeric17/inventory-system/common/responses"
)

type JSONError struct {
	responses.EncodeJSON
	Message string
}

func (jsonErr JSONError) Print(w io.Writer, err error) {
	jsonErr.Message = err.Error()
	responses.ToJSON(w, jsonErr)
}
