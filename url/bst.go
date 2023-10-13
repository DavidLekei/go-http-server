package url

type BSTNode struct {
	data any
	left *BSTNode
	right *BSTNode
}

func BSTCreate(data any) *BSTNode {
	return &BSTNode{
		data: data,
		left: nil,
		right: nil,
	}
}