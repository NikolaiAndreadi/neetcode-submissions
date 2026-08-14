type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack {
		data: make([]int, 0),
	}
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack)  Push(val int) {
	s.data = append(s.data, val)
}

func (s *Stack) Pop() {
	s.data = s.data[:len(s.data)-1]
}

func (s *Stack) Top() int {
	return s.data[len(s.data)-1]
}

type MinStack struct {
	dataStack *Stack
	minStack *Stack
}

func Constructor() MinStack {
	return MinStack {
		dataStack: NewStack(),
		minStack: NewStack(),
	}
}

func (this *MinStack) Push(val int) {
	this.dataStack.Push(val)
	if this.minStack.IsEmpty(){
		this.minStack.Push(val)
		return
	}
	oldMin := this.minStack.Top()
	if oldMin <= val {
		this.minStack.Push(oldMin)
	} else {
		this.minStack.Push(val)
	}

}

func (this *MinStack) Pop() {
	this.dataStack.Pop()
	this.minStack.Pop()
}

func (this *MinStack) Top() int {
	return this.dataStack.Top()
}

func (this *MinStack) GetMin() int {
	return this.minStack.Top()
}
