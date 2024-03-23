package parser

type Node struct {
	Id           string
	Value        interface{}
	NextNodes    []*Node
	PreviousNode *Node
}

func NewNode(id string, value interface{}) *Node {
	return &Node{Id: id, Value: value, NextNodes: make([]*Node, 0)}
}

func (node *Node) AddNextNode(childNode *Node) {
	nodes := append(node.NextNodes, childNode)
	childNode.PreviousNode = node
	node.NextNodes = nodes
}
