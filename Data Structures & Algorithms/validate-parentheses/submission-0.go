type Stack[T any] interface {
	IsEmpty() bool 
	Size() int 

	Push(value T)
	Pop() T
	Peek() T
}
type stack[T any] struct {
	elements []T
}
func NewStack[T any](elements []T) Stack[T] {
	return &stack[T]{elements : elements}
}

func (s *stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

func (s *stack[T]) IsEmpty() bool {
	return len(s.elements) ==0
}

func (s * stack[T]) Pop() T {
	if s.IsEmpty(){
		var zeroValue T
		return zeroValue
	}
	topIndex := len(s.elements) -1
	item := s.elements[topIndex]

	s.elements = s.elements[:topIndex]

	return item
}

func (s *stack[T]) Peek() T {
	if s.IsEmpty(){
		var zeroValue T
		return zeroValue
	}
	return s.elements[len(s.elements)-1]
}

func (s *stack[T]) Size() int {
	return len(s.elements)
}

func isValid(s string) bool {
	stack := NewStack[rune]([]rune{})

	for _, char := range s{
		if char == '(' || char == '[' || char == '{'{
			stack.Push(char)
		}else if char == ')' || char == ']' || char == '}' {
			if stack.IsEmpty() {
				return false
			}

			top := stack.Peek()

			if 	( char == ')' && top !='(' ) ||
				( char == ']' && top !='[' ) ||
				( char == '}' && top !='{' ) {
					return false
				}
			
			_ = stack.Pop()
		}
	}
	return stack.IsEmpty()
}
