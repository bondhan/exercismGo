package reverse

//Reverse ...
func Reverse(s string) string {
	if len(s) < 2 {
		return s
	}

	r := []rune(s)

	reverse := ""
	for i := len(r); i > 0; i-- {
		reverse += string(r[i-1])
	}

	return reverse
}
