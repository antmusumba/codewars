package utils

 
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	// a prime number has no square root
	for x:= 2 ; x*x <= n ; x++ {
		// if you divide a prime number by any other number apart from 1 and itself there will always be a remainder
		if n  % x == 0 {
			return false
		} 
	}
	return true
}

