package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	var prev *NodeL
	current := l.Head

	for current != nil {
		if current.Data == data_ref {
			if prev == nil {
				l.Head = current.Next
			} else {
				prev.Next = current.Next
			}

			if current == l.Tail {
				l.Tail = prev
			}

			current = current.Next
		} else {
			prev = current
			current = current.Next
		}
	}
}
