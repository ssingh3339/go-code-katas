package linkedlist

// ListNode represents a node in a singly linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// NewListNode creates a new list node with the given value
func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

// DoublyListNode represents a node in a doubly linked list
type DoublyListNode struct {
	Val  int
	Prev *DoublyListNode
	Next *DoublyListNode
}

// NewDoublyListNode creates a new doubly linked list node with the given value
func NewDoublyListNode(val int) *DoublyListNode {
	return &DoublyListNode{Val: val}
}

// ArrayToList converts an array to a linked list
func ArrayToList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := NewListNode(arr[0])
	current := head
	for i := 1; i < len(arr); i++ {
		current.Next = NewListNode(arr[i])
		current = current.Next
	}
	return head
}

// ListToArray converts a linked list to an array
func ListToArray(head *ListNode) []int {
	result := []int{}
	current := head
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}

// Length returns the length of the linked list
func Length(head *ListNode) int {
	count := 0
	current := head
	for current != nil {
		count++
		current = current.Next
	}
	return count
}

// Reverse reverses a linked list in-place
func Reverse(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}
