package genericcomparable

import "testing"

func TestComparable(t *testing.T) {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[int8]float64{
		-128: 35.98,
		127:  26.99,
	}

	t.Logf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))

	t.Log(whoIsMin[int64](100, 1000))
}
