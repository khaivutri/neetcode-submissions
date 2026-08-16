type Stack[T any] interface {
	Push(val T)
	Size() int 
	IsEmpty() bool
	Pop()
	Top() T
}

type stack[T any] struct {
	stack []T
}

func NewStack[T any]() Stack[T] {
	return &stack[T]{
		stack: []T{},
	}
}

func (s *stack[T]) Push(val T) {
	s.stack = append(s.stack, val)
}

func (s *stack[T]) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *stack[T]) Pop() {
	if s.IsEmpty() {
		return
	}

	topIndex := len(s.stack) -1

	s.stack = s.stack[:topIndex]
}

func (s *stack[T]) Top() T {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue
	}
	topIndex := len(s.stack) -1
	return s.stack[topIndex]
	
}

func (s *stack[T]) Size() int {
	return len(s.stack)
}
func dailyTemperatures(temperatures []int) []int {
	s := NewStack[int]()
	result := make([]int, len(temperatures))
	for i := range temperatures{
		for !s.IsEmpty() && temperatures[i] > temperatures[s.Top()] {
            top := s.Top()
            s.Pop()
            result[top] = i - top
        }
        s.Push(i)
    }
	return result
}
