package flake

import (
	"testing"
)

func BenchmarkGenerateFlakeID(b *testing.B) {
	// Initialize the FlakeGenerator
	fg, _ := NewGenerator(
		GeneratorWithRegionID(1),
		GeneratorWithMachineID(1),
	)

	// Run the benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fg.GenerateFlakeID()
	}
}
