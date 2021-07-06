func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	maxTotal := (maxChoosableInteger * (maxChoosableInteger + 1)) / 2

	if desiredTotal > maxTotal {
		return false
	}

	if desiredTotal <= maxChoosableInteger {
		return true
	}

	if desiredTotal+maxChoosableInteger < maxTotal {
		return false
	}

	return true
}
