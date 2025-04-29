package service

func Fib(n int) int64 {
	if n == 1 || n == 2 {
		return 1
	} else {
		var a, b int64 = 1,1 
		for i:= 3; i <= n; i++ {
			a, b = b, a+b
		}
		return b
	}
}