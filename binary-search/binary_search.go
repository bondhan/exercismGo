package binarysearch

// SearchInts ...
func SearchInts(slice []int, key int) int {

	half := len(slice) / 2

	isFound := false
	for !isFound {

		if key == slice[half] {
			return half
		} else if key > slice[half] {
			half = (half + len(slice)) / 2
		} else if key < slice[half] {
			half = half / 2
		}

		if (half == 0 || half == len(slice)-1) && key != slice[half] {
			break
		}
	}

	return -1
}
