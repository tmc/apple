package espresso

import "math"

// Argmax returns the index of the maximum value in the slice.
func Argmax(data []float32) int {
	if len(data) == 0 {
		return -1
	}
	best := 0
	max := float32(math.Inf(-1))
	for i, v := range data {
		if v > max {
			max = v
			best = i
		}
	}
	return best
}

// Softmax applies softmax normalization to a float32 slice in-place and returns it.
func Softmax(data []float32) []float32 {
	if len(data) == 0 {
		return data
	}
	max := float32(math.Inf(-1))
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	var sum float64
	for i, v := range data {
		data[i] = float32(math.Exp(float64(v - max)))
		sum += float64(data[i])
	}
	for i := range data {
		data[i] /= float32(sum)
	}
	return data
}

// TopK returns the indices of the K largest values in descending order.
func TopK(data []float32, k int) []int {
	if k <= 0 || len(data) == 0 {
		return nil
	}
	if k > len(data) {
		k = len(data)
	}
	// Simple selection sort for small k.
	indices := make([]int, k)
	used := make([]bool, len(data))
	for j := range k {
		best := -1
		max := float32(math.Inf(-1))
		for i, v := range data {
			if !used[i] && v > max {
				max = v
				best = i
			}
		}
		indices[j] = best
		used[best] = true
	}
	return indices
}
