package main

import (
	"encoding/json"
	"github.com/dawidhermann/json-to-go/parser"
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("input.json")
	if err != nil {
		log.Fatal("Cannot read input.json content", err.Error())
	}
	data := make(map[string]interface{})
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatal("Failed to parse JSON")
	}
	parser.CreateDependencyGraph(data)
	//fmt.Println(reflect.TypeOf(graph.NextNodes[2].NextNodes[0].NextNodes[0].NextNodes[0].NextNodes[0].Value).Kind())
}
