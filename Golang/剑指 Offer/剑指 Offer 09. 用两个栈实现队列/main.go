package main

type CQueue struct {
	stackIn  []int
	stackOut []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.stackIn = append(this.stackIn, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.stackOut) == 0 {
		if len(this.stackIn) == 0 {
			return -1
		}
		for len(this.stackIn) > 0 {
			index := len(this.stackIn) - 1
			this.stackOut = append(this.stackOut, this.stackIn[index])
			this.stackIn = this.stackIn[:index]
		}
	}
	index := len(this.stackOut) - 1
	value := this.stackOut[index]
	this.stackOut = this.stackOut[:index]
	return value
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
