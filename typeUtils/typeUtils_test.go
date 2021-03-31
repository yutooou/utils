package typeUtils

import (
	"testing"
)
// goos: darwin
// goarch: amd64
// pkg: utils/typeUtils
// cpu: Intel(R) Core(TM) i3-7100U CPU @ 2.40GHz
// BenchmarkStringToBytes-4                1000000000               0.4359 ns/op
// BenchmarkStringToBytesWithoutFunc-4     100000000               10.10 ns/op
func BenchmarkStringToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str := "wish you good luck"
		bytes := StringToBytes(str)
		func (bs []byte) {}(bytes)
	}
}
func BenchmarkStringToBytesWithoutFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str := "wish you good luck"
		bytes := []byte(str)
		func (bs []byte) {}(bytes)
	}
}
// goos: darwin
// goarch: amd64
// pkg: utils/typeUtils
// cpu: Intel(R) Core(TM) i3-7100U CPU @ 2.40GHz
// BenchmarkBytesToString-4                1000000000               0.4219 ns/op
// BenchmarkBytesToStringWithoutFunc-4     100000000               10.31 ns/op
func BenchmarkBytesToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes := []byte{119, 105, 115, 104, 32,
						121, 111, 117, 32, 103,
						111, 111, 100, 32, 108,
						117, 99, 107}
		str := BytesToString(bytes)
		func (s string) {}(str)
	}
}
func BenchmarkBytesToStringWithoutFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes := []byte{119, 105, 115, 104, 32,
			121, 111, 117, 32, 103,
			111, 111, 100, 32, 108,
			117, 99, 107}
		str := string(bytes)
		func (s string) {}(str)
	}
}