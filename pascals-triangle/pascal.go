package pascal

func calcTriangle(memory [][]int, n int, max int) [][]int {
	var inner []int
	inner = append(inner, 1)

	if n > 1 {
		mem := memory[n-2]
		for i := 1; i < len(mem); i++ {
			inner = append(inner, mem[i-1]+mem[i])
		}
		inner = append(inner, 1)
	}

	memory = append(memory, inner)

	if n == max {
		return memory
	}
	return calcTriangle(memory, n+1, max)
}

// Triangle ...
func Triangle(n int) [][]int {
	var memory [][]int

	memory = calcTriangle(memory, 1, n)

	return memory
}
