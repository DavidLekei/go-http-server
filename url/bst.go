package url

type BSTNode struct {
	data any
	isVariable bool
	left *BSTNode
	right *BSTNode
}

func BSTCreate(data any, isVariable bool) *BSTNode {
	return &BSTNode{
		data: data,
		isVariable: isVariable,
		left: nil,
		right: nil,
	}
}