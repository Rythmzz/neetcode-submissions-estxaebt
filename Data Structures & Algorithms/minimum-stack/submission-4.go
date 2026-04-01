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
	} else {
		if val <= this.sm[len(this.sm)-1] {
			this.sm = append(this.sm,val)
		}
	}
}

func (this *MinStack) Pop() {
	if  this.s[len(this.s)-1] == this.sm[len(this.sm)-1] {
		this.sm = this.sm[:len(this.sm)-1]
	}

	this.s = this.s[:len(this.s)-1]
}

func (this *MinStack) Top() int {
	top := this.s[len(this.s)-1]
	return top
}

func (this *MinStack) GetMin() int {
	minVal := this.sm[len(this.sm)-1]
	return  minVal
}
