type Stack[T any] interface {
	IsEmpty() bool 
	Size() int

	Push(val T)
	Top() T
	Pop() T
}

type stack[T any] struct {
	elements []T
}

func NewStack[T any]() Stack[T]{
	return &stack[T]{
		elements: []T{},
	}
} 

func (s *stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *stack[T]) Size() int {
	return len(s.elements)
}

func (s *stack[T]) Push(val T) {
	s.elements = append(s.elements, val)

}

func (s *stack[T]) Pop() T{
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue
	}

	topIndex := len(s.elements)-1
	val := s.elements[topIndex]
	s.elements = s.elements[:topIndex]

	return val
}

func (s *stack[T]) Top() T {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue
	}

	return s.elements[len(s.elements)-1]
}

func evalRPN(tokens []string) int {
	stack := NewStack[int]()

	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil {
			stack.Push(val)
		}else{
			s2 := stack.Pop()
			s1 := stack.Pop()

			switch token {
				case "+":
					stack.Push(s2+s1)
				case "-":
					stack.Push(s1-s2)
				case "*":
					stack.Push(s2*s1)
				case "/":
					stack.Push(s1/s2)
			}
		}
	}
	return stack.Pop()
}
