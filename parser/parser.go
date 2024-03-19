// Converts JSON data to go struct
package parser

import (
	"reflect"
	"unsafe"
)

func CreateDependencyGraph(data map[string]interface{}) []*Node {
	for k, v := range data {
		if reflect.ValueOf(v).Kind() == reflect.Map {
			m := *(*map[string]interface{})(unsafe.Pointer(&v))
			return CreateDependencyGraph(m)
		}
	}
}
