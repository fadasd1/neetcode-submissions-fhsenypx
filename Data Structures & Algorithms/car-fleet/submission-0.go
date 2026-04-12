

func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	array := make([][]float64, len(position))
	for i := 0; i < n; i++ {
        array[i] = make([]float64, 2)
		array[i][0] = float64(position[i])
		array[i][1] = float64((target - position[i])) / float64(speed[i])
	}
	sort.Slice(array, func(i, j int) bool {
		return array[j][0] < array[i][0]
	})

	fleets := 0
	lastTime := 0.0

	for i := 0; i < n; i++ {
		time := array[i][1]

		if time > lastTime {
			fleets++
			lastTime = time
		}
	}

	return fleets
}