package typeUtils

import (
	"unsafe"
)
// Converting bytes to string without copying underlying data
// 
func BytesToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}
