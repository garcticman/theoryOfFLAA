package RegularExpressions

type StateIterator struct {
	int32
}

func (si *StateIterator) GetNewState() int32 {
	si.int32++
	return si.int32
}
