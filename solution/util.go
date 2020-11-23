package solution

// ListNode Data structure
type ListNode struct {
	Val  int
	Next *ListNode
}

func makeLinkedList(nums []int) *ListNode {
	dummy := &ListNode{}
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		node := &ListNode{Val: nums[i]}
		node.Next = dummy.Next
		dummy.Next = node
	}
	return dummy.Next
}

// TreeNode Data structure
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
