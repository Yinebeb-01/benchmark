package main

import "testing"

func BenchmarkRegEx(b *testing.B) {
	b.ReportAllocs()
	passwd := "ab05@113zA"
	for n := 0; n < b.N; n++ {
		_ = regEx(passwd)
	}
}

func BenchmarkStringContainsAny(b *testing.B) {
	b.ReportAllocs()
	passwd := "ab05@113zA"
	for n := 0; n < b.N; n++ {
		_ = strContains(passwd)
	}
}
