package maxsubarray

func FindMaxSubarray(arr []int) (subarr []int, sum int) {
	var max, currMax subarray
	for i, x := range arr {
		currMax.len++
		currMax.sum += x
		if x > currMax.sum {
			currMax = subarray{start: i, len: 1, sum: x}
		}
		if currMax.sum > max.sum {
			max = currMax
		}
	}
	return arr[max.start : max.start+max.len], max.sum
}

type subarray struct {
	start, len, sum int
}
