type Pair struct {
	pos 	 	int
	speed 		int
}

func carFleet(target int, position []int, speed []int) int {
	pairs := make([]Pair, len(position))
	for i := range position {
		pairs[i] = Pair{
			pos: position[i],
			speed: speed[i],
		}
	}

	sort.Slice(pairs, func (i, j int) bool {
		if pairs[i].pos == pairs[j].pos {
			return pairs[i].speed > pairs[j].speed
		}
		return pairs[i].pos > pairs[j].pos
	})

	// len, capacity
	stack := make([]float64, 0)
	for i := range pairs {
		time := float64(target - pairs[i].pos) / float64(pairs[i].speed)

		if len(stack) != 0 {
			top := stack[len(stack)-1]

			if time <= top {
				continue
			}
		}
		stack = append(stack, time)
	}
	return len(stack)
	
}
