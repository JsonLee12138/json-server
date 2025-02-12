package core

import (
	"encoding/json"

	"github.com/mailru/easyjson"
)

func MarshalForFiber(v any) ([]byte, error) {
	if v, ok := v.(easyjson.Marshaler); ok {
		return easyjson.Marshal(v)
	}
	return json.Marshal(v)
}

func UnmarshalForFiber(data []byte, v any) error {
	var dest any
	if u, ok := v.(easyjson.Unmarshaler); ok {
		dest = easyjson.Unmarshal(data, u)
	} else {
		dest = json.Unmarshal(data, v)
	}
	return Validator.Struct(dest)
}
