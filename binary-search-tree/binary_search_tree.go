package binarysearchtree

// SearchTreeData ...
type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

// // SearchTreeData ...
// type searchTreeData interface {
// 	Insert(int)
// 	MapString(func(int) string) []string
// 	MapInt(func(int) int) []int
// }

// Bst ...
func Bst(number int) *SearchTreeData {
	return &SearchTreeData{
		left:  nil,
		data:  number,
		right: nil,
	}
}

// Insert ...
func (s *SearchTreeData) Insert(number int) {
	// loop through all the branch from the left to the right
	if number <= s.data {
		if s.left == nil {
			s.left = &SearchTreeData{nil, number, nil}
			return
		}
		s.left.Insert(number)
	} else if number >= s.data {
		if s.right == nil {
			s.right = &SearchTreeData{nil, number, nil}
			return
		}
		s.right.Insert(number)
	}
}

// MapString ...
func (s *SearchTreeData) MapString(f func(int) string) []string {
	var res []string
	if s.left != nil {
		res = append(res, s.left.MapString(f)...)
	}

	res = append(res, f(s.data))

	if s.right != nil {
		res = append(res, s.right.MapString(f)...)
	}

	return res
}

// MapInt ...
func (s *SearchTreeData) MapInt(f func(int) int) []int {
	var res []int
	if s.left != nil {
		res = append(res, s.left.MapInt(f)...)
	}

	res = append(res, f(s.data))

	if s.right != nil {
		res = append(res, s.right.MapInt(f)...)
	}

	return res
}
