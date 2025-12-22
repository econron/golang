package stringrecipes

import (
	"strings"
	"testing"
)

func BenchmarkStringConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := ""
		for j := 0; j < 1000; j++ {
			result += "a"
		}
		_ = result
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for j := 0; j < 1000; j++ {
			builder.WriteString("a")
		}
		_ = builder.String()
	}
}

/*
go test -bench=. -benchmem -benchtime=3s ./stringrecipes
ベンチ結果

BenchmarkStringConcat-8            56664             78280 ns/op          530281 B/op         999 allocs/op
BenchmarkStringBuilder-8         1319317              2776 ns/op            3320 B/op           9 allocs/op

stringsBuilderの方が約28倍速く、約160倍メモリ効率が良く、アロケーション回数が1/111
*/
