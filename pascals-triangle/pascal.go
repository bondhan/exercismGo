package pascal

func calcTriangle(memory [][]int, n int) [][]int {
	var inner []int
	inner = append(inner, 1)

	if n > 1 {
		mem := memory[n-2]
		for i := 1; i < len(mem); i++ {
			inner = append(inner, mem[i-1]+mem[i])
		}
		inner = append(inner, 1)
	}

	return append(memory, inner)
}

func Triangle(n int) [][]int {
	var memory [][]int

	for i := 1; i <= n; i++ {
		memory = calcTriangle(memory, i)
	}

	return memory
}
