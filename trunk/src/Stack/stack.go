// Stack project Stack.go
package stack

type Stack struct {
	i    int
	data [10]int
}

func (s *Stack) Push(val int) {
	s.data[s.i] = val
	s.i++
}

func (s *Stack) Pop() (val int) {
	s.i--
	val = s.data[s.i]
	return
}
