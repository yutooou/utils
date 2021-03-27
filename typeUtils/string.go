package typeUtils

import (
	"unsafe"
)
// Converting string to bytes without copying underlying data
// Note: since string is immutable, do not change the content of bytes
func StringToBytes(str string) []byte {
	ps := (*[2]uintptr)(unsafe.Pointer(&str))
	// reflect.SliceHeader
	bp := [3]uintptr{ps[0], ps[1], ps[1]}
	return *(*[]byte)(unsafe.Pointer(&bp))
}
