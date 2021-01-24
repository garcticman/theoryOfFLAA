package PDA

type PDAConfiguration struct {
	state int32
	stack Stack
}

type PDARule struct {
	state          int32
	character      int32
	nextState      int32
	popCharacter   int32
	pushCharacters []int32
}

func (rule *PDARule) AppliesTo(configuration PDAConfiguration, character int32) bool {
	return rule.state == configuration.state &&
		rule.popCharacter == configuration.stack.Top() &&
		rule.character == character
}
func (rule *PDARule) Follow(configuration PDAConfiguration) PDAConfiguration {
	return PDAConfiguration{
		state: rule.nextState,
		stack: rule.NextStack(configuration),
	}
}
func (rule PDARule) NextStack(configuration PDAConfiguration) Stack {
	poppedStack := configuration.stack.Pop()

	reversed := reverseArray(rule.pushCharacters)
	for _, v := range reversed {
		poppedStack = poppedStack.Push(v)
	}

	return poppedStack
}

func reverseArray(array []int32) []int32 {
	lenx := len(array)                    // lenx holds the original array length
	reversed_array := make([]int32, lenx) // creates a slice that refer to a new array of length lenx

	for i := 0; i < lenx; i++ {
		j := lenx - (i + 1) // j initially holds (lenx - 1) and decreases to 0 while i initially holds 0 and increase to (lenx - 1)
		reversed_array[i] = array[j]
	}

	return reversed_array
}
