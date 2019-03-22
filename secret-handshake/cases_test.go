package secret

// Source: exercism/problem-specifications
// Commit: a9e4df8 secret-handshake: Apply new "input" policy
// Problem Specifications Version: 1.2.0

type secretHandshakeTest struct {
	code uint
	h    []string
}

var tests = []secretHandshakeTest{
	{1, []string{"wink"}},
	{2, []string{"double blink"}},
	{4, []string{"close your eyes"}},
	{8, []string{"jump"}},
	{3, []string{"wink", "double blink"}},
	{19, []string{"double blink", "wink"}},
	{24, []string{"jump"}},
	{16, []string{}},
	{15, []string{"wink", "double blink", "close your eyes", "jump"}},
	{31, []string{"jump", "close your eyes", "double blink", "wink"}},
	{0, []string{}},
}

var consts = []string{"", "wink", "double blink", "close your eyes", "jump"}

func Handshake(code uint) []string {
	var compile []string

	var index uint
	for index = 1; index <= 16; index++ {
		if (code & index) == uint(1) {
			compile = append(compile, consts[1])
		} else if (code & index) == uint(2) {
			compile = append(compile, consts[2])
		} else if (code & index) == uint(4) {
			compile = append(compile, consts[3])
		} else if (code & index) == uint(8) {
			compile = append(compile, consts[3])
		}
	}

	return compile
}
