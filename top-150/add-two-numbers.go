/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1CurrentNode := l1
	l2CurrentNode := l2
	remainder := 0
	var previousNode *ListNode
	var topNode *ListNode
	for {
		l1NodeValue := 0
		if l1CurrentNode == nil {

		} else {
			l1NodeValue = l1CurrentNode.Val
			l1CurrentNode = l1CurrentNode.Next
		}

		l2NodeValue := 0
		if l2CurrentNode == nil {
		} else {
			l2NodeValue = l2CurrentNode.Val
			l2CurrentNode = l2CurrentNode.Next
		}

		sum := l1NodeValue + l2NodeValue + remainder
		remainder = sum / 10
		currentNode := &ListNode{
			Val: sum % 10,
		}
		if topNode == nil {
			topNode = currentNode
			previousNode = currentNode
		} else {
			previousNode.Next = currentNode
			previousNode = currentNode
		}
		if remainder == 0 && l1CurrentNode == nil && l2CurrentNode == nil {
			return topNode
		}
	}
}
