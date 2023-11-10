package sliding_window_maximum

import "fmt"

type Dequeue struct {
	first *Node
	last  *Node
}

type Node struct {
	next  *Node
	prev  *Node
	value int
}

func (d *Dequeue) push(elem int) {
	node := &Node{value: elem}

	if d.isEmpty() {
		d.first = node
		d.last = d.first
		return
	}

	node.prev = d.last
	d.last.next = node
	d.last = node
}

func (d *Dequeue) pop() int {
	value := d.last.value
	d.last = d.last.prev

	if d.last == nil {
		d.first = nil
	} else {
		d.last.next = nil
	}

	return value
}

func (d *Dequeue) popFront() int {
	value := d.first.value
	d.first = d.first.next

	if d.first == nil {
		d.last = nil
	} else {
		d.first.prev = nil
	}

	return value
}

func (d *Dequeue) isEmpty() bool {
	return d.first == nil && d.last == nil
}

func (d *Dequeue) back() int {
	return d.last.value
}

func (d *Dequeue) front() int {
	return d.first.value
}

func (d *Dequeue) print() {
	curr := d.first
	for curr != nil {
		fmt.Printf("%d ", curr.value)
		curr = curr.next
	}
}

func MaxSlidingWindow(nums []int, k int) []int {
	return maxSlidingWindow(nums, k)
}

func maxSlidingWindow(nums []int, k int) []int {
	maxNumsInWindows := make([]int, 0, len(nums)-k+1)

	indexDq := Dequeue{first: nil, last: nil}

	for i := 0; i < k; i++ {
		for !indexDq.isEmpty() && nums[i] >= nums[indexDq.back()] {
			indexDq.pop()
		}
		indexDq.push(i)
	}

	maxNumsInWindows = append(maxNumsInWindows, nums[indexDq.front()])
	for i := k; i < len(nums); i++ {
		if i-indexDq.front() >= k {
			indexDq.popFront()
		}

		for !indexDq.isEmpty() && nums[i] >= nums[indexDq.back()] {
			indexDq.pop()
		}

		indexDq.push(i)
		maxNumsInWindows = append(maxNumsInWindows, nums[indexDq.front()])
	}

	return maxNumsInWindows
}
