type MinStack struct {
	elements []int
	min 	 []int
}

func Constructor() MinStack {
	return MinStack{
		elements: []int{},
		min:	  []int{},
	}
}

func (this * MinStack) IsEmpty(mode string) bool{
	switch mode {
		case "elements":
			return len(this.elements) == 0
		case "min":
			return len(this.min) == 0
		default:
			return false
	}
}
func (this *MinStack) Push(val int) {
	if this.IsEmpty("elements") || val <= this.GetMin() {
		this.min = append(this.min, val)
	}
	this.elements = append(this.elements, val)
}

func (this *MinStack) Pop() {
	if this.IsEmpty("elements") {
		return
	}

	topIndex := len(this.elements) - 1
	val := this.elements[topIndex]
	this.elements = this.elements[:topIndex]

	if val <= this.GetMin() {
		this.min = this.min[:len(this.min)-1]
	}
}

func (this *MinStack) Top() int {	
	if this.IsEmpty("elements") {
		return 0
	}
	return this.elements[len(this.elements)-1]
}

func (this *MinStack) GetMin() int {
	if this.IsEmpty("min") {
		return 0
	}

	return this.min[len(this.min)-1]
}
