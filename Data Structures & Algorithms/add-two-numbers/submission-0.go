/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    stackl1 := []int{}
	stackl2 := []int{}

	for l1 != nil {
		stackl1 = append(stackl1,l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		stackl2 = append(stackl2,l2.Val)
		l2 = l2.Next
	}

	num1, num2 := 0,0
	for len(stackl1) > 0 {
		num := stackl1[len(stackl1)-1]
		num1 = num1*10 + num
		stackl1 = stackl1[:len(stackl1)-1]
	}

	for len(stackl2) > 0 {
		num := stackl2[len(stackl2)-1]
		num2 = num2*10 + num
		stackl2 = stackl2[:len(stackl2)-1]
	}

	sum := num1 + num2

	fmt.Printf("num1 %d, num2 %d, sum %d",num1,num2,sum)

	nodeNew := &ListNode{}
	curr := nodeNew

	for sum != 0 {
		curr.Val = sum%10
		sum = sum/10
		if sum != 0 {
		curr.Next = &ListNode{}
		}
		curr = curr.Next
	}

	return nodeNew


	

}
