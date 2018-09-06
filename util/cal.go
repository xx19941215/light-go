package util

func Cal(n1 float64, n2 float64, operate byte) float64 {
	var res float64
	switch operate {
		case '+':
		res = n1 + n2
		case '-':
		res = n1 - n2
		case '*':
		res = n1 * n2
		case '/':
		res = n1 / n2
	}

	return res
}