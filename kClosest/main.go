// есть плоскость x и y
// необходимо найти k - количество ближайших точек к 0 и вернуть их координаты

package main

import (
	"container/heap"
	"fmt"
	"math"
	"math/rand/v2"
)

type point struct {
	x, y int
}

var (
	points = randomPoints(1000)
	k      = 50
)

func randomPoints(n int) []point {
	pts := make([]point, n)
	for i := range n {
		x := rand.IntN(2*n) - n
		y := rand.IntN(2*n) - n
		p := point{x, y}
		pts[i] = p
	}
	return pts
}

func main() {
	fmt.Println(SearchClosest(points, k))
	fmt.Println(SearchClosestOptimized(points, k))
}

func SearchClosest(data []point, k int) []point {
	ds := make([]float64, 0, len(data))
	for _, p := range data {
		d := math.Sqrt(float64(p.x*p.x + p.y*p.y))
		ds = append(ds, d)
	}

	for i := 0; i < k; i++ {
		imin := i
		for j := i + 1; j < len(ds); j++ {
			if ds[j] < ds[imin] {
				imin = j
			}
		}
		data[i], data[imin] = data[imin], data[i]
		ds[i], ds[imin] = ds[imin], ds[i]
	}
	return data[:k]
}

func SearchClosestOptimized(data []point, k int) []point {
	type pointD struct {
		point
		d float64
	}

	dataD := make([]pointD, len(data))
	for i, p := range data {
		dataD[i].point = p
		dataD[i].d = math.Sqrt(float64(p.x*p.x + p.y*p.y))
	}

	for i := 0; i < k; i++ {
		imin := i
		for j := i + 1; j < len(dataD); j++ {
			if dataD[j].d < dataD[imin].d {
				imin = j
			}
		}
		data[i], data[imin] = data[imin], data[i]
		dataD[i], dataD[imin] = dataD[imin], dataD[i]
	}
	return data[:k]
}

func dist2(p point) int {
	return p.x*p.x + p.y*p.y
}

/* ---------- Max Heap ---------- */

type item struct {
	p point
	d int // squared distance
}

type MaxHeap []item

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].d > h[j].d } // max-heap
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(item))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

/* ---------- Solution ---------- */

func SearchClosestHeap(data []point, k int) []point {
	if k <= 0 || k > len(data) {
		return nil
	}

	h := &MaxHeap{}
	heap.Init(h)

	for _, p := range data {
		d := dist2(p)
		it := item{p: p, d: d}

		if h.Len() < k {
			heap.Push(h, it)
		} else if d < (*h)[0].d {
			heap.Pop(h)
			heap.Push(h, it)
		}
	}

	res := make([]point, h.Len())
	for i := 0; h.Len() > 0; i++ {
		res[i] = heap.Pop(h).(item).p
	}
	return res
}
