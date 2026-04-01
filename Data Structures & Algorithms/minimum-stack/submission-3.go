type MinStack struct {
	st []int
	sm []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.st  = append(this.st,val)

	if len(this.sm) == 0 || this.sm[len(this.sm)-1] >= val {
		this.sm = append(this.sm,val)
	}

}

func (this *MinStack) Pop() {
	if this.sm[len(this.sm)-1] == this.st[len(this.st)-1] {
		this.sm = this.sm[:len(this.sm)-1]
	}

	this.st = this.st[:len(this.st)-1]
}

func (this *MinStack) Top() int {
	return this.st[len(this.st)-1]
}

func (this *MinStack) GetMin() int {
	return this.sm[len(this.sm)-1]
}
