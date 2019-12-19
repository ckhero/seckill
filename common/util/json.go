package util

import "encoding/json"

func JsonEncode(val interface{}) string  {
	if val == nil {
		return "{}"
	}
	str, err := json.Marshal(val)
	if err != nil {
		return "{}"
	}
	return string(str);
}
