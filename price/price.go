package price

func GetRoundOffAmount(totalPriceCent int64) int64 {
	lastDigit := totalPriceCent % 10

	switch lastDigit {
	case 1, 2:
		//If the figure ends in 1,2 then the rounding off figure will go down
		return -lastDigit
	case 3, 4, 6, 7:
		//If the figure ends in 3,4 then the rounding off figure will go up
		//If the figure ends in 6,7 then the rounding off figure will go down
		return 5 - lastDigit
	case 8, 9:
		//If the figure ends in 8,9 then the rounding off figure will go up
		return 10 - lastDigit
	default:
		return 0
	}
}
