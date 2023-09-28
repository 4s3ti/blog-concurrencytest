package main

import "testing"

const (
	file100 = "data/systems_100.json"
	file1000 = "data/systems_1000.json"
	file10000 = "data/systems_10000.json"
	file100000 = "data/systems_100000.json"
	file1000000 = "data/systems_1000000.json"
)

// Benchmarking
func BenchmarkNameSearch(b *testing.B){
	b.Run("100 Systems", func (b *testing.B) {
		names := []string{"system0", "system50", "system99", "foo"}
		for i := 0; i < b.N; i++ {
			_, _ = SystemsByName(file100, names)
		}
	})

	b.Run("1000 Systems", func (b *testing.B) {
		names := []string{"system0", "system500", "system999", "foo"}
		for i := 0; i < b.N; i++ {
			_, _ = SystemsByName(file1000, names)
		}
	})

	b.Run("10000 Systems", func (b *testing.B) {
		names := []string{"system0", "system5000", "system9999", "foo"}
		for i := 0; i < b.N; i++ {
			_, _ = SystemsByName(file10000, names)
		}
	})

	b.Run("100000 Systems", func (b *testing.B) {
		names := []string{"system0", "system50000", "system99999", "foo"}
		for i := 0; i < b.N; i++ {
			_, _ = SystemsByName(file100000, names)
		}
	})

	b.Run("1000000 Systems", func (b *testing.B) {
		names := []string{"system0", "system50000", "system999999", "foo"}
		for i := 0; i < b.N; i++ {
			_, _ = SystemsByName(file1000000, names)
		}
	})
}
