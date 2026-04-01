type MinStack struct {
	store []int
	minStore []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.store = append(this.store,val)
	if len(this.minStore) == 0 || this.minStore[len(this.minStore)-1] >= val {
		this.minStore = append(this.minStore,val)
	}
}

func (this *MinStack) Pop() {
	if len(this.store) == 0 {
		 return
	}

	if len(this.minStore) > 0 && this.minStore[len(this.minStore)-1] == this.store[len(this.store)-1] {
		this.minStore = this.minStore[:len(this.minStore)-1]
	}

	this.store = this.store[:len(this.store)-1]

}

func (this *MinStack) Top() int {
	topValue := this.store[len(this.store)-1]
	return topValue
}

func (this *MinStack) GetMin() int {	
	return this.minStore[len(this.minStore)-1]
}
