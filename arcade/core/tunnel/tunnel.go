package tunnel

func leastFactorial(n int) int {
	k := 1
	for i := 1; k < n; i++ {
		k *= i
	}
	return k
}

func countSumOfTwoRepresentations2(n int, l int, r int) int {
	cnt := 0
	a := n / 2
	b := n - a
	for l <= a && b <= r {
		if a+b == n {
			cnt++
		}
		a--
		b++
	}
	return cnt
}

//func magicalWell(a int, b int, n int) int {
//	return a*b*n + (n*n*(a+b)-a*n-b*n)/2 + (n*(n-1)*(2*n-1))/6
//}

func magicalWell(a int, b int, n int) int {
	s := 0
	for ; n > 0; n-- {
		s += a * b
		a++
		b++
	}
	return s
}

func countBlackCells(n int, m int) int {
	cnt := 0
	slope := float64(n) / float64(m)
	for x := 0.0; x < float64(m); x++ {
		for y := 0.0; y < float64(n); y++ {
			v := x * slope
			if v >= y && v <= (y+1) {
				cnt++
			} else if (y <= (v + slope)) && ((v + slope) <= (y + 1)) {
				cnt++
			}
		}
	}
	return cnt
}
