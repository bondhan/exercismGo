package binarysearch

// SearchInts ...
func SearchInts(slice []int, key int) int {

	if len(slice) < 1 {
		return -1
	}

	low := 0
	high := len(slice) - 1
	half := (low + high) / 2

	for true {

		//check if it is at the middle
		if key == slice[half] {
			return half
		} else if key > slice[half] { // if key larger than the middle, make starting point from the half
			low = half + 1
		} else if key < slice[half] { // if key is smaller, make end point is the half
			high = half - 1
		}

		half = (low + high) / 2 //compute the new half

		if low > high {
			break
		}
	}

	return -1
}
