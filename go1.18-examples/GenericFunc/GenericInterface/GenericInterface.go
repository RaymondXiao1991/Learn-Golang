package genericinterface

type Element interface {
	int64 | float64 | string
}

type Stack[V Element] struct {
	size  int
	value []V
}

func (s *Stack[V]) Push(v V) {
	s.value = append(s.value, v)
	s.size++
}

func (s *Stack[V]) Pop() V {
	n := s.size - 1
	e := s.value[n]
	if s.size != 0 {
		s.value = s.value[:n]
		s.size--
	}
	return e
}
