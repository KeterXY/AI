package optimizer

type stack struct {
	body	 interface{}
	isBottom bool
	next	 *stack
}

func generateStack() *stack {
	op := new(stack)
	op.isBottom = true
	op.next = nil
	op.body = nil
	return op
}

func(srcStack *stack)push(input interface{}) *stack {
	newStack := new(stack)
	newStack.isBottom = false
	newStack.body = input
	newStack.next = srcStack
	return newStack
}

func(srcStack *stack)pop() (*stack,interface{}) {
	if srcStack.isBottom == true{
		return nil, nil
	}else {
		return srcStack.next, srcStack.body
	}
}
