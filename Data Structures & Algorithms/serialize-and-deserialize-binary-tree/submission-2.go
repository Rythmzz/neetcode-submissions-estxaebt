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
    result := []int{}

	for _, d := range decode {
		num, _:= strconv.Atoi(d)
		result = append(result,num)
	}

	root := this.buildTree(&result)

	return root
}

func (this *Codec) buildTree(result *[]int) *TreeNode {
	if len(*result) == 0 {
		return nil
	}

	val := (*result)[0]
	*result = (*result)[1:]

	if val == 1001 {
		return nil
	}

	curr := &TreeNode{Val: val}
	curr.Left = this.buildTree(result)
	curr.Right = this.buildTree(result)

	return curr
}
