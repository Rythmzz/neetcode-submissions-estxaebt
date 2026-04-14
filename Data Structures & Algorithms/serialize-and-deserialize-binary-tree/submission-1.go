/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
     var s strings.Builder
	 this.encode(root,&s)
	 return s.String()
}

func (this *Codec) encode(curr *TreeNode, s *strings.Builder) {
	if curr == nil {
		s.WriteString("1001,")
		return
	}

	s.WriteString(strconv.Itoa(curr.Val))
	s.WriteString(",")

	this.encode(curr.Left,s)
	this.encode(curr.Right,s)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
    decode := strings.Split(data,",")
	list := []int{}

	for _,d := range decode {
		num, _:= strconv.Atoi(d)
		list = append(list,num)
	}

	root := this.buildTree(&list)

	return root

}

func (this *Codec) buildTree(list *[]int) *TreeNode {
	if len(*list) == 0 {
		return nil
	}

	val := (*list)[0]
	*list = (*list)[1:]

	if val == 1001 {
		return nil
	}

	curr := &TreeNode{Val: val}

	curr.Left = this.buildTree(list)
	curr.Right = this.buildTree(list)

	return curr
}
