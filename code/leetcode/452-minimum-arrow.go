package main

import "sort"

func findMinArrowShots(points [][]int) int {
	// Edge case: no balloons to burstbtw
	if len(points) == 0 {
		return 0
	}

	// Step 1: Sort balloons by their end position (xEnd)
	// Why? Process most-constrained balloons first (those ending earliest)
	sort.Slice(points, func(i, j int) bool {
		// Use int64 to avoid overflow when comparing endpoints
		return int64(points[i][1]) < int64(points[j][1])
	})

	// Step 2: Initialize tracking variables
	arrows := 1                         // We'll shoot at least one arrow
	lastArrowPos := int64(points[0][1]) // Shoot first arrow at first balloon's end

	// Step 3: Iterate through remaining balloons
	for i := 1; i < len(points); i++ {
		// Check if current balloon overlaps with last arrow position
		// Balloon is burst if: xStart <= lastArrowPos
		// Balloon is NOT burst if: xStart > lastArrowPos
		if int64(points[i][0]) > lastArrowPos {
			// No overlap! Must shoot a new arrow at this balloon's end
			arrows++
			lastArrowPos = int64(points[i][1])
		}
		// If there IS overlap, balloon is already burst by lastArrowPos—skip it
	}

	return arrows
}
