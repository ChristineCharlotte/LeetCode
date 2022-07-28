package main

func subtractProductAndSum(n int) int {
	a := n / 10000
	b := (n - 10000*a) / 1000
	c := (n - 10000*a - 1000*b) / 100
	d := (n - 10000*a - 1000*b - 100*c) / 10
	e := n - 10000*a - 1000*b - c*100 - 10*d
	// fmt.Println(a, b, c, d)
	re := 0
	if a == 0 {
		if b == 0 {
			if c == 0 {
				if d == 0 {
					re = 0
				} else {
					re = d*e - (d + e)
				}
			} else {
				re = c*d*e - (c + d + e)
			}
		} else {
			re = b*c*d*e - (b + c + d + e)
		}

	} else {
		re = a*b*c*d*e - (a + b + c + d + e)
	}
	return re
}
