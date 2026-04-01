type MinStack struct {
	store []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.store = append(this.store, val)
}

func (this *MinStack) Pop() {
	if len(this.store) > 0 {
		this.store = this.store[:len(this.store)-1]
	}
}

func (this *MinStack) Top() int {
	top := this.store[len(this.store)-1]
	return top
}

func (this *MinStack) GetMin() int {
	minNum := this.store[0]
	for i := 0; i < len(this.store); i++ {
		minNum = min(minNum, this.store[i])
	}

	return minNum
}