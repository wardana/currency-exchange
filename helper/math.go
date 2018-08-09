package helper

//MinMaxAverageInSlices get minimum maximum and average in array slice
func (h *Helper) MinMaxAverageInSlices(array []float64) (float64, float64, float64) {
	var min = array[0]
	var max, count, average float64

	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
		count += value
	}
	arrLen := float64(len(array))
	average = count / arrLen
	return min, max, average
}
