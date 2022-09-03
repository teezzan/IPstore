package ipstore

import "testing"

var ipList []string
var ipLookupSize = 1000000000 //change accordingly

func init() {
	ipList = generateIPList(ipLookupSize)
}
func BenchmarkRequestHandled(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		RequestHandled(ipList[n%len(ipList)])
	}
}

func BenchmarkTop100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Top100()
	}
}
