package utils

type Stack []interface{}

func (stack *Stack) IsEmpty() bool {
	return len(*stack) == 0
}

func (stack *Stack) Push(data interface{}) {
	*stack = append(*stack, data)
}

func (stack *Stack) Pop() (interface{}, bool) {
	if stack.IsEmpty() {
		return struct{}{}, false
	} else {
		index := len(*stack) - 1
		data := (*stack)[index]
		*stack = (*stack)[:index]
		return data, true
	}
}

func (stack *Stack) Last() interface{} {
	return (*stack)[len(*stack)-1]
}
