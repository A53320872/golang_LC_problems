package main

func largestRectangleArea(heights []int) int {
	heights = append(heights, -1)
	maxArea := 0
	monotonic := [][]int{}
	i := 0
	for i < len(heights) {
		if len(monotonic) == 0 || heights[i] >= monotonic[len(monotonic)-1][1] {
			monotonic = append(monotonic, []int{i, heights[i]})
		} else {
			leftBound := i
			topH := monotonic[len(monotonic)-1][1]
			for len(monotonic) != 0 && topH >= heights[i] {
				topI := monotonic[len(monotonic)-1][0]
				topH = monotonic[len(monotonic)-1][1]
				if topH >= heights[i] {
					if topH*(i-topI) > maxArea {
						maxArea = topH * (i - topI)
					}
					monotonic = monotonic[:len(monotonic)-1]
					leftBound = topI
				}
			}
			monotonic = append(monotonic, []int{leftBound, heights[i]})
		}
		i++
	}

	return maxArea
}

func main() {

}
