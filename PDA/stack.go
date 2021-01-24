package PDA

type Stack struct {
	contents []int32
}

func (s Stack) Push(character int32) Stack {
	s.contents = append([]int32{character}, s.contents...)
	return s
}
func (s Stack) Pop() Stack {
	if len(s.contents) <= 0 {
		return s
	}

	s.contents = s.contents[1:]
	return s
}
func (s Stack) Top() int32 {
	if len(s.contents) <= 0 {
		return -1
	}
	return s.contents[0]
}
