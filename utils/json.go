package utils

import "encoding/json"

func Marshal(v *interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func UnMarshal(v *interface{}, raw []byte) error {
	return json.Unmarshal(raw, v)
}
