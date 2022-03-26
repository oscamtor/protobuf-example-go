package main

import (
	"fmt"

	example_simple "github.com/oscamtor/protobuf-example-go/src/simple"
)

func main() {
	//sm :=
	doSimple()
}

func doSimple() *example_simple.SimpleMessage {
	sm := example_simple.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 7, 8},
	}
	fmt.Println(sm)

	sm.Name = "I renamed you"
	fmt.Println(sm)

	fmt.Println("The ID is:", sm.GetId())

	return &sm
}
