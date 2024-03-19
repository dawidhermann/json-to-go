package parser

type Node struct {
	Id           string
	Value        interface{}
	NextNodes    []*Node
	PreviousNode *Node
}

func New(id string, value interface{}) *Node {
	return &Node{Id: id}
}

func (node *Node) AddNextNode(childNode *Node) {
	nodes := append(node.NextNodes, childNode)
	node.NextNodes = nodes
}
