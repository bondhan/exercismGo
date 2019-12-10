package binarysearch

// SearchInts ...
func SearchInts(slice []int, key int) int {

	l := 0
	r := len(slice) - 1

	for m := (l + r) / 2; l <= r; m = (l + r) / 2 {

		//check if it is at the middle
		if key == slice[m] {
			return m
		} else if key > slice[m] { // if key larger than the middle, make starting point from the half
			l = m + 1
		} else if key < slice[m] { // if key is smaller, make end point is the half
			r = m - 1
		}
	}

	return -1
}
