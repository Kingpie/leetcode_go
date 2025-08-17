package leetcode_go

/*
Implement a first in first out (FIFO) queue using only two stacks. The implemented queue should support all the functions of a normal queue (push, peek, pop, and empty).

Implement the MyQueue class:

void push(int x) Pushes element x to the back of the queue.
int pop() Removes the element from the front of the queue and returns it.
int peek() Returns the element at the front of the queue.
boolean empty() Returns true if the queue is empty, false otherwise.
Notes:

You must use only standard operations of a stack, which means only push to top, peek/pop from top, size, and is empty operations are valid.
Depending on your language, the stack may not be supported natively. You may simulate a stack using a list or deque (double-ended queue) as long as you use only a stack's standard operations.


Example 1:

Input
["MyQueue", "push", "push", "peek", "pop", "empty"]
[[], [1], [2], [], [], []]
Output
[null, null, null, 1, 1, false]

Explanation
MyQueue myQueue = new MyQueue();
myQueue.push(1); // queue is: [1]
myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
myQueue.peek(); // return 1
myQueue.pop(); // return 1, queue is [2]
myQueue.empty(); // return false
*/

type MyQueue struct {
	In, Out []int
}

func (this *MyQueue) in2out() {
	for len(this.In) > 0 {
		this.Out = append(this.Out, this.In[len(this.In)-1])
		this.In = this.In[:len(this.In)-1]
	}
}

func Constructor() MyQueue {
	return MyQueue{
		In:  make([]int, 0),
		Out: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.In = append(this.In, x)
}

func (this *MyQueue) Pop() int {
	if len(this.Out) == 0 {
		this.in2out()
	}
	x := this.Out[len(this.Out)-1]
	this.Out = this.Out[:len(this.Out)-1]
	return x
}

func (this *MyQueue) Peek() int {
	if len(this.Out) == 0 {
		this.in2out()
	}
	return this.Out[len(this.Out)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.In) == 0 && len(this.Out) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
