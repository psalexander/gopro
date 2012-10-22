package stack

import "testing"

func TestPush(t *testing.T) {
	s := new(Stack)
	testNum := 10
	s.Push(testNum)
	if testNum != s.Pop() {
		t.Log("get the wrong num!")
		t.Fail()
	}
}
