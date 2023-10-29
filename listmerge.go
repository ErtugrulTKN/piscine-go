package piscine

func ListMerge(l1 *List, l2 *List) {
	if l2.Head != nil {
		if l1.Tail != nil {
			l1.Tail.Next = l2.Head
			l1.Tail = l2.Tail
		} else {
			l1.Head = l2.Head
			l1.Tail = l2.Tail
		}
	}
	l2.Head = nil
	l2.Tail = nil
}
