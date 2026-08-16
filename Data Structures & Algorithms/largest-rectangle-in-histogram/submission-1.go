func largestRectangleArea(heights []int) int {
    n := len(heights)
    // Cấp phát trước bộ nhớ để tránh reallocate
    stack := make([]int, 0, n+1)
    maxArea := 0

    // Duyệt từ 0 đến n (với i == n đại diện cho cột sentinel chiều cao = 0)
    for i := 0; i <= n; i++ {
        curHeight := 0
        if i < n {
            curHeight = heights[i]
        }

        // Duy trì stack đơn điệu tăng dần
        for len(stack) > 0 && heights[stack[len(stack)-1]] > curHeight {
            // Pop đỉnh stack
            topIdx := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            h := heights[topIdx]

            // Xác định chiều rộng
            width := i
            if len(stack) > 0 {
                width = i - stack[len(stack)-1] - 1
            }

            if area := h * width; area > maxArea {
                maxArea = area
            }
        }

        stack = append(stack, i)
    }

    return maxArea
}