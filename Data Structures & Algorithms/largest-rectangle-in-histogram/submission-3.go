type Pairs struct {
	i int
	h int
}
func largestRectangleArea(heights []int) int {
	stack := make([]Pairs, 0)
	n := len(heights)

	maxArea :=0
	for i, h := range heights {
		start := i
		for len(stack) > 0 && stack[len(stack)-1].h > h {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			
			prevIdx, prevH := top.i, top.h
			maxArea = max(maxArea, prevH*(i-prevIdx))

            start = prevIdx 
		}
		stack = append(stack, Pairs{
			i: start,
			h: h,
		})
	}

	for _, pair := range stack {
		maxArea = max(maxArea, pair.h*(n-pair.i))
	}
	
	return maxArea
}
