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
    var encode strings.Builder
	this.encodeString(root,&encode)
	fmt.Println(encode.String())
	return encode.String()
}

func (this *Codec) encodeString(curr *TreeNode, encode *strings.Builder) {
		if curr == nil {
			encode.WriteString(",0")
			return
		}
		st := fmt.Sprintf(",%d",curr.Val)

		if encode.String() == "" {
			st = fmt.Sprintf("%d",curr.Val)
		}
		
		encode.WriteString(st)

		this.encodeString(curr.Left,encode)
		this.encodeString(curr.Right,encode)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	queueStr := strings.Split(data,",")
	queue := []int{}
	for _, s := range queueStr {
		num,_ := strconv.Atoi(s)
		queue = append(queue,num)
	}
	fmt.Println(queue)
	curr := this.buildTree(&queue)
    return curr
}

func (this *Codec) buildTree(queue *[]int) *TreeNode {
	if len(*queue) == 0 {
		return nil
	}

	val := (*queue)[0]
	*queue = (*queue)[1:]

	if val == 0 {
		return nil
	}

	curr := &TreeNode{Val:val}

	curr.Left = this.buildTree(queue)
	curr.Right = this.buildTree(queue)

	return curr
}
