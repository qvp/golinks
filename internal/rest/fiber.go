package rest

import (
	"encoding/json"
	"reflect"
)

// JSONEncoder returns [] then slice or array is nil
func JSONEncoder(v interface{}) ([]byte, error) {
	if v == nil { // todo better solution ?
		typ := reflect.TypeOf(v).Kind()
		if typ == reflect.Array || typ == reflect.Slice {
			return []byte("[]"), nil
		}
	}
	return json.Marshal(v)
}
