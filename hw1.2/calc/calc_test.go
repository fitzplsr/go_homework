package calc

import (
	"errors"
	"fmt"
	"testing"
)

var testsSuccess = []struct {
	in  string
	out float64
}{
	{in: "3 + 4 * 2 / ( 1 - 5 ) --2--3",
		out: 6},

	{in: "3+4*2/(1-5)--2--3",
		out: 6},

	{in: "2 + 2*2",
		out: 6},

	{in: "2.0+2.0*2.0",
		out: 6},

	{in: "(2.0+2.0)*2.0",
		out: 8},

	{in: "-1",
		out: -1},

	{in: "-2--2",
		out: 0},
}

func TestCalcSuccess(t *testing.T) {
	for n, tt := range testsSuccess {
		t.Run(fmt.Sprintf("TEST#%d", n), func(t *testing.T) {
			res, _ := Calc(tt.in)
			if res != tt.out {
				t.Errorf("got %g, want %g", res, tt.out)
			}
		})
	}
}

var testsFailureIncorrectExpressionError = []struct {
	in  string
	err *IncorrectExpressionError
}{
	{in: "1 + 1()",
		err: &IncorrectExpressionError{}},

	{in: "(-1",
		err: &IncorrectExpressionError{}},

	{in: "1+1)",
		err: &IncorrectExpressionError{}},
}

func TestCalcFailureIncorrectExpression(t *testing.T) {
	for n, tt := range testsFailureIncorrectExpressionError {
		t.Run(fmt.Sprintf("TEST#%d", n), func(t *testing.T) {
			_, err := Calc(tt.in)
			if !errors.As(err, &tt.err) {
				t.Errorf("Calc don't throw IncorrectExpressionError on data:%s", tt.in)
			}
		})
	}
}

var testsFailureComputeError = []struct {
	in  string
	err *ComputeError
}{
	{in: "(1) (2)",
		err: &ComputeError{}},

	{in: "1 / ( 1-1 )",
		err: &ComputeError{}},

	{in: "-",
		err: &ComputeError{}},
}

func TestCalcFailureCompute(t *testing.T) {
	for n, tt := range testsFailureComputeError {
		t.Run(fmt.Sprintf("TEST#%d", n), func(t *testing.T) {
			_, err := Calc(tt.in)
			if !errors.As(err, &tt.err) {
				t.Errorf("Calc don't throw ComputeError on data:%s", tt.in)
			}
		})
	}
}
