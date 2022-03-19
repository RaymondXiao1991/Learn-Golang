package genericinterface

import "testing"

func TestStack(t *testing.T) {
	// INT STACK
	strS := Stack[int64]{}
	strS.Push(1)
	strS.Push(2)
	strS.Push(3)
	t.Log(strS.Pop())
	t.Log(strS.Pop())
	t.Log(strS.Pop())

	// FLOAT STACK
	floatS := Stack[float64]{}
	floatS.Push(1.1)
	floatS.Push(2.2)
	floatS.Push(3.3)
	t.Log(floatS.Pop())
	t.Log(floatS.Pop())
	t.Log(floatS.Pop())
}
