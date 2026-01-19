package main

import "testing"

// run in console
// go test -bench . -benchmem

func BenchmarkSearchClosest(b *testing.B) {
	for b.Loop() {
		SearchClosest(points, k)
	}
}

func BenchmarkSearchClosestOptimized(b *testing.B) {
	for b.Loop() {
		SearchClosestOptimized(points, k)
	}
}

func BenchmarkSearchClosestHeap(b *testing.B) {
	for b.Loop() {
		SearchClosestHeap(points, k)
	}
}
