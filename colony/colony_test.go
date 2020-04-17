package colony

import (
	"testing"
)

func BenchmarkSeeder(b *testing.B) {
	game := New(200, 60)

	b.ReportAllocs()
	b.ResetTimer()

	for x := 0; x < b.N; x++ {
		game.Seed()
	}
}

func BenchmarkIncubation(b *testing.B) {
	game := New(200, 60)

	b.ReportAllocs()
	b.ResetTimer()

	for x := 0; x < b.N; x++ {
		game.Incubate()
	}
}
