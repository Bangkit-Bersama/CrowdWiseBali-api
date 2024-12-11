package main

import (
	"fmt"

	"github.com/galeone/tfgo"
)

func main() {
	model := tfgo.LoadModel("./models/model_occupancy_fix", []string{"serve"}, nil)
	fmt.Println(model)
}
