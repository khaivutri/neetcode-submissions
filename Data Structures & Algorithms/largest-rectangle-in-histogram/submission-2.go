type Pairs struct {
    index  int
    height int
}

func largestRectangleArea(heights []int) int {
    stack := make([]Pairs, 0, len(heights))
    maxArea := 0
    n := len(heights)

    for i, h := range heights {
        start := i
        for len(stack) > 0 && stack[len(stack)-1].height > h {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1] 

            prevIdx, prevH := top.index, top.height
            maxArea = max(maxArea, prevH*(i-prevIdx))

            start = prevIdx 
        }
        stack = append(stack, Pairs{index: start, height: h})
    }

    for _, pair := range stack {
        maxArea = max(maxArea, pair.height*(n-pair.index))
    }

    return maxArea
}