package deposit

//Calculate func
func Calculate(amount int, currency string) (min int, max int) {
	percentMin, percentMax := 0,0
	if currency == "TJS" {
		percentMin, percentMax = 4,6
	} 
	if currency == "USD" {
		percentMin, percentMax = 1,2
	}
	min = (amount * percentMin / 100) / 12
	max = (amount * percentMax / 100) / 12
	return min,max
}
