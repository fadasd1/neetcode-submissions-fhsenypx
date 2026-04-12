type Car struct {
	pos  int
	time float64
}

func carFleet(target int, position []int, speed []int) int {

	n := len(position)
	array := make([]Car, len(position))
	for i := 0; i < n; i++ {
		array[i] = Car{
			pos:  position[i],
			time: float64((target - position[i])) / float64(speed[i]),
		}

	}
	sort.Slice(array, func(i, j int) bool {
		return array[j].pos < array[i].pos
	})

	fleets := 0
	lastTime := 0.0

	for i := 0; i < n; i++ {
		time := array[i].time

		if time > lastTime {
			fleets++
			lastTime = time
		}
	}

	return fleets
}