package main

type MinStack struct {
	source     []int
	minElement []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		source:     []int{},
		minElement: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.source = append(this.source, x)
	if len(this.minElement) == 0 || this.minElement[len(this.minElement)-1] >= x {
		this.minElement = append(this.minElement, x)
	}
}

func (this *MinStack) Pop() {
	temp := this.source[len(this.source)-1]
	this.source = this.source[:len(this.source)-1]
	if this.minElement[len(this.minElement)-1] == temp {
		this.minElement = this.minElement[:len(this.minElement)-1]
	}
}

func (this *MinStack) Top() int {
	return this.source[len(this.source)-1]
}

func (this *MinStack) Min() int {
	return this.minElement[len(this.minElement)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
