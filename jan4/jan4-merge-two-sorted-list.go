package jan4

// ListNode Singly linkedlist
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// var result *ListNode
	// for {
	//     if l1 != nil && l2 != nil{

	//         if l1.Val <= l2.Val{
	//             result.Next = l1
	//             l1 = l1.Next
	//         }else{
	//             result.Next = l2
	//             l2 = l2.Next
	//         }
	//     }else{
	//         break
	//     }
	//     result= result.Next
	// }
	//     if l1 != nil {
	// 	result.Next = l1
	//     }
	//     if l2 != nil {
	//        result.Next = l2
	//     }

	// return result
	var a *ListNode
	pos := &ListNode{Val: -1, Next: nil}
	a = pos
	for {
		if l1 != nil && l2 != nil {

			if l1.Val <= l2.Val {
				pos.Next = l1
				l1 = l1.Next
			} else {
				pos.Next = l2
				l2 = l2.Next
			}
			pos = pos.Next
		} else {
			break
		}

	}
	if l1 != nil {
		pos.Next = l1
	}
	if l2 != nil {
		pos.Next = l2
	}

	return a.Next
}
