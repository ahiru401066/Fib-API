package service

import (
	"testing"
	"math/big"
)

func TestFib(t *testing.T) {
	testcase := []struct {
		name           string
		input          int
		expectedOutput *big.Int
		expectedError  bool
	}{
		{"n=10", 10, big.NewInt(55), false},
		{"n=2", 2, big.NewInt(1), false},
		{"n=1", 1, big.NewInt(1), false},
		{"n=0", 0, big.NewInt(0), true},
		{"n=-1", -1, big.NewInt(0), true},
		{"n=-2", -2, big.NewInt(0), true},
		{"n=-5", -5, big.NewInt(0), true},
	}

	for _, tc := range testcase {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Fib(tc.input)
			if tc.expectedError {
				if err == nil {
					t.Errorf("エラーを期待したが、err == nil")
				}
				return
			}

			if err != nil {
				t.Error("エラーなしを期待したが、err")
				return
			}

			if result.Cmp(tc.expectedOutput) != 0 {
				t.Error("計算結果が一致しない")
			}
		})
	}
}
