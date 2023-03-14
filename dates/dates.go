package dates

const DaysPerWeek int = 7

func WeeksToDays(weeks int) int {
	return weeks * DaysPerWeek
}

func DaysToWeeks(days int) float64 {
	return float64(days) / float64(DaysPerWeek)
}
