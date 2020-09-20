package transfer

// Total func
func Total(amount int) (bonus int) {
	Percent := 5
	bonus = amount * Percent / 1000
	return bonus + amount
}