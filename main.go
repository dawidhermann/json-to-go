package jsontogo

import (
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("input.json")
	if err != nil {
		log.Fatal("Cannot read input.json content", err.Error())
	}
	
}