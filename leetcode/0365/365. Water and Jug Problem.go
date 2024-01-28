package _365

type jugState struct {
	x, y int
}

func canMeasureWater(jug1Capacity, jug2Capacity, targetCapacity int) bool {
	stack := []jugState{{0, 0}}
	seen := make(map[jugState]struct{})

	for len(stack) > 0 {
		state := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if state.x == targetCapacity || state.y == targetCapacity || state.x+state.y == targetCapacity {
			return true
		}
		if _, exists := seen[state]; exists {
			continue
		}
		seen[state] = struct{}{}

		// Fill jug1
		stack = append(stack, jugState{jug1Capacity, state.y})
		// Fill jug2
		stack = append(stack, jugState{state.x, jug2Capacity})
		// Empty jug1
		stack = append(stack, jugState{0, state.y})
		// Empty jug2
		stack = append(stack, jugState{state.x, 0})
		// Pour jug1 into jug2
		stack = append(stack, jugState{state.x - min(state.x, jug2Capacity-state.y), state.y + min(state.x, jug2Capacity-state.y)})
		// Pour jug2 into jug1
		stack = append(stack, jugState{state.x + min(state.y, jug1Capacity-state.x), state.y - min(state.y, jug1Capacity-state.x)})
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
