package daily_temperatures

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

func New() *Stack {
	return &Stack{nil, 0}
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) Top() interface{} {
	if s.length == 0 {
		return nil
	}

	return s.top.value
}

func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}

	n := s.top
	s.top = n.prev
	s.length--

	return n.value
}

func (s *Stack) Push(value interface{}) {
	n := &node{value, s.top}
	s.top = n
	s.length++
}

func DailyTemperatures(temperatures []int) []int {
	return dailyTemperatures(temperatures)
}

// second solution with Stack struct, just for fun
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	nextTemperatures := make([]int, n)

	stack := New()

	for i := n - 1; i >= 0; i-- {
		if stack.Len() == 0 {
			stack.Push(i)
			nextTemperatures[i] = 0
			continue
		}

		for temperatures[i] >= temperatures[stack.Top().(int)] {
			stack.Pop()

			if stack.Len() == 0 {
				break
			}
		}

		if stack.Len() == 0 {
			nextTemperatures[i] = 0
		} else {
			nextTemperatures[i] = stack.Top().(int) - i
		}
		stack.Push(i)
	}

	return nextTemperatures
}
