package service

import (
	"fmt"
	"math/big"
) 

func Fib(n int) (*big.Int ,error) {
	if n <= 0 {
		return new(big.Int), fmt.Errorf("bad param") // 0とエラーを返す
	}
	if n == 1 || n == 2 {
		return big.NewInt(1),nil
	} else {
		a := big.NewInt(1) 
		b := big.NewInt(1) 
		for i:= 3; i <= n; i++ {
			next := new(big.Int).Add(a,b)
			a, b = b, next
		}
		return b, nil
	}
}