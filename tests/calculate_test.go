package tests

import (
	"GoCalc/calculator"
	"testing"
)

func TestForArrayExpressions(t *testing.T) {
	expressions := []string{"2+2", "2+2*2", "(2+2)*2", "56+(13*2-(7)+23/5-(59+(2*56-(13*5))-25*12)+50)-9"}
	result := []float64{4.0, 6.0, 8.0, 314.6}

	for i := 0; i < len(expressions); i++ {
		got, _ := calculator.Calculate(expressions[i])
		want := result[i]
		if got != want {
			t.Errorf("got %f, wanted %f", got, want)
		}
	}
}
