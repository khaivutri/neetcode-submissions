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

func (this *MinStack) IsEmpty() bool{
	return len(this.elements) ==0
}
func (this *MinStack) Push(val int) {
	this.elements = append(this.elements, val)

	if len(this.min) ==0 || val <= this.GetMin() {
		this.min = append(this.min, val)
	}
}

func (this *MinStack) Pop() {
	if this.IsEmpty() {
		return
	}

	topIndex := len(this.elements) -1
	val := this.elements[topIndex]

	if val <= this.GetMin() {
		this.min = this.min[:len(this.min)-1]	
	}
	this.elements = this.elements[:topIndex]
}

func (this *MinStack) Top() int {
	if this.IsEmpty() {
		return 0
	}
	
	return this.elements[len(this.elements)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.min) == 0 {
		return 0
	}
	return this.min[len(this.min)-1]
	

}
