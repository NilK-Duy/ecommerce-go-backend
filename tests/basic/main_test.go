package basic

import (
	"testing"
)

func TestAddOne(t *testing.T) {
	var (
		input  = 1
		output = 3
	)

	actual := AddOne(1)
	if actual != output {
		t.Errorf("AddOne(%d), input %d, actual %d", input, output, actual)
	}
}

func TestAddOne2(t *testing.T) {
	var (
		input  = 1
		output = 3
	)

	actual := AddOne2(1)
	if actual != output {
		t.Errorf("AddOne(%d), input %d, actual %d", input, output, actual)
	}
}

// func TestRequire(t *testing.T) {
// 	require.Equal(t, 2, 3)
// 	fmt.Println("Not executing")
// }

// func TestAssert(t *testing.T) {
// 	assert.Equal(t, 2, 3)
// 	fmt.Println("executing")
// }
