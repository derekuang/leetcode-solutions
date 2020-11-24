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

func makeTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{Val: nums[0]}
	nodes := make([]*TreeNode, len(nums))
	nodes[0] = root
	for i := 1; i < len(nums); i++ {
		parent := nodes[(i-1)/2]
		if nums[i] == -1 {
			continue
		} else {
			nodes[i] = &TreeNode{Val: nums[i]}
			if i%2 == 1 {
				parent.Left = nodes[i]
			} else {
				parent.Right = nodes[i]
			}
		}
	}
	return root
}
