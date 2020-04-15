package game

import (
	"testing"
)

func BenchmarkGame(b *testing.B) {
	game := NewGame(200, 60)

	b.SetBytes(200 * 60 * 4)
	b.ReportAllocs()
	b.ResetTimer()

	for x := 0; x < b.N; x++ {
		game.Incubate()
	}
}
