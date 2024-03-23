// Package parser converts JSON data to go struct
package parser

import (
	"fmt"
	"log"
)

func CreateDependencyGraph(data map[string]interface{}) *Node {
	typesMap := make(map[string][]*Node)
	rootNode := NewNode("root", nil)
	res := createDependencyGraph(rootNode, data, typesMap)
	//if val, ok := typesMap["bool"]; ok {
	fmt.Println(typesMap["query_string"])
	//}
	return res
}

func createDependencyGraph(rootNode *Node, data map[string]interface{}, typesMap map[string][]*Node) *Node {
	for k, v := range data {
		switch val := v.(type) {
		case map[string]interface{}:
			node := createDependencyGraph(NewNode(k, nil), val, typesMap)
			rootNode.AddNextNode(node)
			typeEntry, ok := typesMap[node.Id]
			if !ok {
				typeEntry = make([]*Node, 1)
			}
			typeEntry = append(typeEntry, node)
			typesMap[node.Id] = typeEntry
		case []interface{}:
			for _, e := range val {
				if value, ok := e.(map[string]interface{}); ok {
					node := createDependencyGraph(NewNode(k, nil), value, typesMap)
					rootNode.AddNextNode(node)
					typeEntry, ok := typesMap[node.Id]
					if !ok {
						typeEntry = []*Node{node}
					} else {
						typeEntry = append(typeEntry, node)
					}
					typesMap[node.Id] = typeEntry
				} else {
					log.Fatal("Not a map")
				}
			}
		default:
			node := NewNode(k, v)
			rootNode.AddNextNode(node)
		}
	}
	return rootNode
}
