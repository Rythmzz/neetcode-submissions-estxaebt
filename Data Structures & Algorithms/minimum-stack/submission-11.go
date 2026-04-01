type MinStack struct {
	s []int
	sm []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.s = append(this.s,val)

	if len(this.sm) == 0 {
		this.sm = append(this.sm,val)
		return
	}

	if this.sm[len(this.sm)-1] >= val {
		this.sm = append(this.sm,val)
	}
}

func (this *MinStack) Pop() {
	if this.s[len(this.s)-1] == this.sm[len(this.sm)-1] {
		this.sm = this.sm[:len(this.sm)-1]
	}

	this.s = this.s[:len(this.s)-1]

}

func (this *MinStack) Top() int {
	return this.s[len(this.s)-1]
}

func (this *MinStack) GetMin() int {
	return this.sm[len(this.sm)-1]
}
