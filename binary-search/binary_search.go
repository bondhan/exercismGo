package binarysearch

// SearchInts ...
func SearchInts(slice []int, key int) int {

	isFound := false
	for !isFound {
		half := len(slice) / 2
		if key == slice[half] {
			return half
		} else if key > slice[half] {
			return half + SearchInts(slice[half:], key)
		} else if key < slice[half] {
			return SearchInts(slice[:half], key)
		}

		if half == 1 {
			break
		}
	}

	return 0
}
