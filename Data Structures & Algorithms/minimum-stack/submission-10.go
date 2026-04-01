type MinStack struct {
	s []int
	m []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.s = append(this.s,val)

	if len(this.m) == 0 {
		this.m = append(this.m,val)
	} else {
		if this.m[len(this.m)-1] >= val {
			this.m = append(this.m,val)
		}
	}
}

func (this *MinStack) Pop() {
	if this.s[len(this.s)-1] == this.m[len(this.m)-1] {
		this.m = this.m[:len(this.m)-1]
	}

	this.s = this.s[:len(this.s)-1]
}

func (this *MinStack) Top() int {
	return this.s[len(this.s)-1]
}

func (this *MinStack) GetMin() int {
	return this.m[len(this.m)-1]
}
