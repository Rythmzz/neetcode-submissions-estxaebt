type MinStack struct {
	store []int
	minStore []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.store = append(this.store,val)

	m := len(this.minStore)
	if m == 0 || this.minStore[m-1] >= val {
		this.minStore =  append(this.minStore,val)
	}
}

func (this *MinStack) Pop() {
	m:= len(this.minStore)
	n:= len(this.store)

	if this.store[n-1] == this.minStore[m-1] {
		this.minStore = this.minStore[:m-1]
	}

	this.store = this.store[:n-1]
}

func (this *MinStack) Top() int {
	n:= len(this.store)
	return this.store[n-1]
}

func (this *MinStack) GetMin() int {
	m:= len(this.minStore)
	return this.minStore[m-1]
}
