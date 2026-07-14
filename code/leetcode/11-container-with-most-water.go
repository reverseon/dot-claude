package main

func maxArea(height []int) int {
	// Two-pointer approach: start at widest possible container.
	// Starting wide maximizes the width component of area (width × min_height).
	// As we shrink inward, we lose width, so we need the height to increase
	// enough to compensate—this only happens by finding taller lines.
	left, right := 0, len(height)-1
	maxArea := 0

	// Keep shrinking the window until pointers converge (no more containers to check).
	for left < right {
		// The area of a container is determined by:
		// - Width: horizontal distance between lines (right - left)
		// - Height: limited by the shorter of the two lines (bottleneck principle)
		// Area = width × min(height[left], height[right])
		width := right - left
		h := height[left]
		if height[right] < h {
			h = height[right]
		}
		area := width * h

		// Track the best container we've found so far.
		if area > maxArea {
			maxArea = area
		}

		// The key decision: which pointer to move?
		// - If we move the TALLER line inward: width shrinks, height is still bottlenecked
		//   by the shorter line, so area will only shrink. No benefit.
		// - If we move the SHORTER line inward: width shrinks (bad), but we might find
		//   a taller line that increases the bottleneck height. This is the ONLY way
		//   area can improve when moving inward. So we must try this path.
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
