package main

func countGoodRectangles(rectangles [][]int) int {
	largestLen := 0
	count := 0
	for _, rect := range rectangles {
		if rect[0] < rect[1] && rect[0] > largestLen {
			largestLen = rect[0]
		} else if rect[1] < rect[0] && rect[1] > largestLen {
			largestLen = rect[1]
		}
	}

	for _, rect := range rectangles {
		if rect[0] >= largestLen && rect[1] >= largestLen {
			count++
		}
	}

	return count
}
