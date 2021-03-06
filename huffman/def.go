package huffman

// TreeNode 树节点
type TreeNode struct {
	Character byte      `json:"character"` //字母
	Weight    int       `json:"weight"`    //权重
	FNode     *TreeNode `json:"fnode"`     //父节点
	LNode     *TreeNode `json:"lnode"`     //左节点
	RNode     *TreeNode `json:"rnode"`     //右节点
	Code      string    `json:"code"`      // 编码
}

// TreeNodeSlice 树节点切片
type TreeNodeSlice []*TreeNode
