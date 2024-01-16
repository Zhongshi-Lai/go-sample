package pkg

import "fmt"

type Tools struct {
}

func NewAllTools() *Tools {

	fmt.Println("start init all tools")

	return &Tools{}
}
