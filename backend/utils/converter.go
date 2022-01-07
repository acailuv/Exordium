package utils

import (
	"bytes"
	"encoding/json"
)

func ConvertToByteArray(data interface{}) []byte {
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(data)

	return buffer.Bytes()
}
