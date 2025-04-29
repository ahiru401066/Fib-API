package utils

func Fib(n int) (int64,error) {
	if n == 1 || n == 2 {
		return 1,nil
	} else {
		var a, b int64 = 1,1 
		for i:= 3; i <= n; i++ {
			a, b = b, a+b
		}
		return b, nil
	}
}